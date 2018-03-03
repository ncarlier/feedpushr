package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/aggregator"
	"github.com/ncarlier/feedpushr/pkg/controller"
	"github.com/ncarlier/feedpushr/pkg/logging"
	"github.com/ncarlier/feedpushr/pkg/metric"
	"github.com/ncarlier/feedpushr/pkg/opml"
	"github.com/ncarlier/feedpushr/pkg/output"
	"github.com/ncarlier/feedpushr/pkg/store"
	"github.com/rs/zerolog/log"
)

var (
	importFilename = flag.String("import", "", "Import a OPML file at boostrap")
	outputURI      = flag.String("output", Config.Output, "Output destination URI")
	clearCache     = flag.Bool("clear-cache", false, "Clear cache at bootstrap")
)

func main() {
	flag.Parse()

	if *version {
		printVersion()
		os.Exit(0)
	}

	// Log configuration
	logging.Configure(Config.LogLevel, Config.LogPretty)

	// Metric configuration
	metric.Configure()

	// Init the data store
	db, err := store.Configure(Config.Store)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init data store")
	}

	// Clear cache if require
	if *clearCache {
		err = db.ClearCache()
		if err != nil {
			log.Fatal().Err(err).Msg("unable to clear the cache")
		}
	}

	// Starts cache-buster
	cleanCacheTicker := time.NewTicker(24 * time.Hour)
	go func() {
		log.Debug().Str("retention", Config.CacheRetention.String()).Msg("cache-buster started")
		for _ = range cleanCacheTicker.C {
			maxAge := time.Now().Add(-Config.CacheRetention)
			err := db.EvictFromCache(maxAge)
			if err != nil {
				log.Error().Err(err).Msg("unable clean the cache")
				return
			}
		}
	}()

	// Init output manager
	om, err := output.NewManager(db, *outputURI, Config.CacheRetention)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init output manager")
	}

	// Import OPML file if present
	if *importFilename != "" {
		o, err := opml.NewOPMLFromFile(*importFilename)
		if err != nil {
			db.Close()
			log.Fatal().Err(err)
		}
		err = opml.ImportOPMLToDB(o, db)
		if err != nil {
			log.Error().Err(err)
		}
	}
	// Init aggregator daemon
	var callbackURL string
	if Config.PublicURL != "" {
		callbackURL = Config.PublicURL + "/v1/pshb"
	}
	am, err := aggregator.NewManager(db, om, Config.Delay, callbackURL)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Create service
	service := goa.New("feedpushr")

	// Set custom logger
	logger := log.With().Str("component", "server").Logger()
	service.WithLogger(logging.NewLogAdapter(logger))

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(false))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "feed" controller
	fc := controller.NewFeedController(service, db, am)
	app.MountFeedController(service, fc)
	// Mount "health" controller
	hc := controller.NewHealthController(service)
	app.MountHealthController(service, hc)
	// Mount "swagger" controller
	sc := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, sc)
	// Mount "opml" controller
	oc := controller.NewOpmlController(service, db)
	app.MountOpmlController(service, oc)
	// Mount "vars" controller
	vc := controller.NewVarsController(service)
	app.MountVarsController(service, vc)
	// Mount "pshb" controller (only if public URL is configured)
	if Config.PublicURL != "" {
		pc := controller.NewPshbController(service, db, am, om)
		app.MountPshbController(service, pc)
	}

	// Graceful shutdown handle
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		log.Debug().Msg("shutting down server...")
		cleanCacheTicker.Stop()
		service.CancelAll()
		service.Server.SetKeepAlivesEnabled(false)
		am.Shutdown()

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := service.Server.Shutdown(ctx); err != nil {
			log.Fatal().Err(err).Msg("could not gracefully shutdown the server")
		}
		close(done)
	}()

	// Start service
	if err := service.ListenAndServe(Config.ListenAddr); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("unable to start server")
	}

	<-done
	db.Close()
	log.Debug().Msg("server stopped")
}
