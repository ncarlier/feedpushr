package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/aggregator"
	"github.com/ncarlier/feedpushr/pkg/assets"
	"github.com/ncarlier/feedpushr/pkg/config"
	"github.com/ncarlier/feedpushr/pkg/controller"
	"github.com/ncarlier/feedpushr/pkg/filter"
	"github.com/ncarlier/feedpushr/pkg/job"
	"github.com/ncarlier/feedpushr/pkg/logging"
	"github.com/ncarlier/feedpushr/pkg/metric"
	"github.com/ncarlier/feedpushr/pkg/opml"
	"github.com/ncarlier/feedpushr/pkg/output"
	"github.com/ncarlier/feedpushr/pkg/plugin"
	"github.com/ncarlier/feedpushr/pkg/store"
	"github.com/rs/zerolog/log"
)

func main() {
	flag.Parse()

	if *config.Version {
		printVersion()
		os.Exit(0)
	}

	// Log configuration
	logging.Configure(*config.LogLevel, *config.LogPretty, config.SentryDSN)

	// Load plugins
	pr, err := plugin.NewPluginRegistry(config.Plugins.Values())
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init plugins")
	}

	// Metric configuration
	metric.Configure()

	// Init the data store
	db, err := store.Configure(*config.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init data store")
	}

	// Clear cache if require
	if *config.ClearCache {
		err = db.ClearCache()
		if err != nil {
			log.Fatal().Err(err).Msg("unable to clear the cache")
		}
	}

	// Starts background jobs (cache-buster)
	scheduler := job.StartNewScheduler(db)

	// Init chain filter
	cf, err := filter.NewChainFilter(config.Filters.Values(), pr)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init filter chain")
	}

	// Init output manager
	om, err := output.NewManager(db, config.Outputs.Values(), *config.CacheRetention, pr, cf)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init output manager")
	}

	// Import OPML file if present
	if *config.ImportFilename != "" {
		o, err := opml.NewOPMLFromFile(*config.ImportFilename)
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
	if *config.PublicURL != "" {
		callbackURL = *config.PublicURL + "/v1/pshb"
	}
	am, err := aggregator.NewManager(db, om, *config.Delay, *config.Timeout, callbackURL)
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
	// Mount "filter" controller
	fic := controller.NewFilterController(service, cf)
	app.MountFilterController(service, fic)
	// Mount "output" controller
	oc := controller.NewOutputController(service, om)
	app.MountOutputController(service, oc)
	// Mount "health" controller
	hc := controller.NewHealthController(service)
	app.MountHealthController(service, hc)
	// Mount "swagger" controller
	sc := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, sc)
	// Mount "opml" controller
	opc := controller.NewOpmlController(service, db)
	app.MountOpmlController(service, opc)
	// Mount "vars" controller
	vc := controller.NewVarsController(service)
	app.MountVarsController(service, vc)
	// Mount "pshb" controller (only if public URL is configured)
	if *config.PublicURL != "" {
		pc := controller.NewPshbController(service, db, am, om)
		app.MountPshbController(service, pc)
	}
	// Mount custom handlers (aka: not generated)...
	service.Mux.Handle("GET", "/ui/*asset", assets.Handler())
	service.Mux.Handle("GET", "/ui/", assets.Handler())

	// Graceful shutdown handle
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		log.Debug().Msg("shutting down server...")
		scheduler.Shutdown()
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
	log.Info().Str("listen", *config.ListenAddr).Msg("starting HTTP server...")
	if err := service.ListenAndServe(*config.ListenAddr); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("unable to start server")
	}

	<-done
	db.Close()
	log.Debug().Msg("server stopped")
}
