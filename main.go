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
	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/controller"
	"github.com/ncarlier/feedpushr/pkg/filter"
	"github.com/ncarlier/feedpushr/pkg/logging"
	"github.com/ncarlier/feedpushr/pkg/metric"
	"github.com/ncarlier/feedpushr/pkg/opml"
	"github.com/ncarlier/feedpushr/pkg/output"
	"github.com/ncarlier/feedpushr/pkg/plugin"
	"github.com/ncarlier/feedpushr/pkg/store"
	"github.com/rs/zerolog/log"
)

var (
	listenAddr,
	dataStore,
	outputURI,
	importFilename,
	publicURL,
	logLevel *string
	delay,
	timeout,
	cacheRetention *time.Duration
	clearCache,
	logPretty *bool
	plugins,
	filters common.ArrayFlags
)

func init() {
	listenAddr = flag.String("addr", Config.ListenAddr, "HTTP server address")
	dataStore = flag.String("db", Config.Store, "Data store location")
	outputURI = flag.String("output", Config.Output, "Output destination")
	importFilename = flag.String("import", "", "Import a OPML file at boostrap")
	publicURL = flag.String("public-url", Config.PublicURL, "Public URL used for PSHB subscriptions")
	delay = flag.Duration("delay", Config.Delay, "Delay between aggregations")
	timeout = flag.Duration("timeout", Config.Timeout, "Aggregation timeout")
	clearCache = flag.Bool("clear-cache", false, "Clear cache at bootstrap")
	cacheRetention = flag.Duration("cache-retention", Config.CacheRetention, "Cache retention duration")
	logPretty = flag.Bool("log-pretty", Config.LogPretty, "Writes log using plain text format")
	logLevel = flag.String("log-level", Config.LogLevel, "Logging level")

	flag.Var(&plugins, "plugin", "Plugin to load")
	filters = Config.Filters
	flag.Var(&filters, "filter", "Filter to apply")
}

func main() {
	flag.Parse()

	if *version {
		printVersion()
		os.Exit(0)
	}

	// Log configuration
	logging.Configure(*logLevel, *logPretty)

	// Load plugins
	pr, err := plugin.NewPluginRegistry(plugins)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init plugins")
	}

	// Metric configuration
	metric.Configure()

	// Init the data store
	db, err := store.Configure(*dataStore)
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
		log.Debug().Str("retention", (*cacheRetention).String()).Msg("cache-buster started")
		for range cleanCacheTicker.C {
			maxAge := time.Now().Add(-*cacheRetention)
			err := db.EvictFromCache(maxAge)
			if err != nil {
				log.Error().Err(err).Msg("unable clean the cache")
				return
			}
		}
	}()

	// Init chain filter
	cf, err := filter.NewChainFilter(filters, pr)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init filter chain")
	}

	// Init output manager
	om, err := output.NewManager(db, *outputURI, *cacheRetention, pr, cf)
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
	if *publicURL != "" {
		callbackURL = *publicURL + "/v1/pshb"
	}
	am, err := aggregator.NewManager(db, om, *delay, *timeout, callbackURL)
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
	if *publicURL != "" {
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
	if err := service.ListenAndServe(*listenAddr); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("unable to start server")
	}

	<-done
	db.Close()
	log.Debug().Msg("server stopped")
}
