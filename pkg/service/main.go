package service

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v2/pkg/assets"
	"github.com/ncarlier/feedpushr/v2/pkg/config"
	"github.com/ncarlier/feedpushr/v2/pkg/controller"
	"github.com/ncarlier/feedpushr/v2/pkg/logging"
	"github.com/ncarlier/feedpushr/v2/pkg/opml"
	"github.com/ncarlier/feedpushr/v2/pkg/pipeline"
	"github.com/ncarlier/feedpushr/v2/pkg/plugin"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
	"github.com/rs/zerolog/log"
)

// Service is the global service
type Service struct {
	db         store.DB
	srv        *goa.Service
	aggregator *aggregator.Manager
}

// ClearCache clear DB cache
func (s *Service) ClearCache() error {
	return s.db.ClearCache()
}

// ImportOPMLFile imports OPML file
func (s *Service) ImportOPMLFile(filename string) error {
	o, err := opml.NewOPMLFromFile(filename)
	if err != nil {
		return err
	}
	err = opml.ImportOPMLToDB(o, s.db)
	if err != nil {
		log.Error().Err(err)
	}
	return nil
}

// ListenAndServe starts server
func (s *Service) ListenAndServe(ListenAddr string) error {
	log.Debug().Msg("loading feed aggregators...")
	if err := loadFeedAggregators(s.db, s.aggregator); err != nil {
		return err
	}
	log.Debug().Msg("starting HTTP server...")
	if err := s.srv.ListenAndServe(ListenAddr); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Shutdown service
func (s *Service) Shutdown(ctx context.Context) error {
	s.aggregator.Shutdown()
	s.srv.CancelAll()
	s.srv.Server.SetKeepAlivesEnabled(false)
	return s.srv.Server.Shutdown(ctx)
}

// Configure the global service
func Configure(db store.DB, conf config.Config) (*Service, error) {
	// Auto add plugins if not configured
	plugins := conf.Plugins
	if len(plugins) == 0 {
		var err error
		plugins, err = filepath.Glob("feedpushr-*.so")
		if err != nil {
			log.Error().Err(err).Msg("unable to autoload plugins")
		}
	}
	// Load plugins
	err := plugin.Configure(plugins)
	if err != nil {
		log.Error().Err(err).Msg("unable to init plugins")
		return nil, err
	}

	// Clear configuration if asked
	if conf.ClearConfig {
		if err := db.ClearFilters(); err != nil {
			log.Error().Err(err).Msg("unable to clear filters")
			return nil, err
		}
		if err := db.ClearOutputs(); err != nil {
			log.Error().Err(err).Msg("unable to clear outputs")
			return nil, err
		}
	}

	// Init chain filter
	cf, err := loadChainFilter(db)
	if err != nil {
		log.Error().Err(err).Msg("unable to init filter chain")
		return nil, err
	}

	// Init the pipeline
	om, err := pipeline.NewPipeline(db, conf.CacheRetention)
	if err != nil {
		log.Error().Err(err).Msg("unable to init output manager")
		return nil, err
	}
	om.ChainFilter = cf

	// Init aggregator daemon
	var callbackURL string
	if conf.PublicURL != "" {
		callbackURL = conf.PublicURL + "/v1/pshb"
	}
	am := aggregator.NewManager(om, conf.Delay, conf.Timeout, callbackURL)

	// Create service
	srv := goa.New("feedpushr")

	// Set custom logger
	logger := log.With().Str("component", "server").Logger()
	srv.WithLogger(logging.NewLogAdapter(logger))

	// Mount middleware
	srv.Use(middleware.RequestID())
	srv.Use(middleware.LogRequest(false))
	srv.Use(middleware.ErrorHandler(srv, true))
	srv.Use(middleware.Recover())

	// Mount "index" controller
	app.MountIndexController(srv, controller.NewIndexController(srv))
	// Mount "feed" controller
	app.MountFeedController(srv, controller.NewFeedController(srv, db, am))
	// Mount "filter" controller
	app.MountFilterController(srv, controller.NewFilterController(srv, db, cf))
	// Mount "output" controller
	app.MountOutputController(srv, controller.NewOutputController(srv, db, om))
	// Mount "health" controller
	app.MountHealthController(srv, controller.NewHealthController(srv))
	// Mount "swagger" controller
	app.MountSwaggerController(srv, controller.NewSwaggerController(srv))
	// Mount "opml" controller
	app.MountOpmlController(srv, controller.NewOpmlController(srv, db))
	// Mount "vars" controller
	app.MountVarsController(srv, controller.NewVarsController(srv))
	// Mount "pshb" controller (only if public URL is configured)
	if conf.PublicURL != "" {
		app.MountPshbController(srv, controller.NewPshbController(srv, db, am, om))
	}
	// Mount custom handlers (aka: not generated)...
	srv.Mux.Handle("GET", "/ui/*asset", assets.Handler())
	srv.Mux.Handle("GET", "/ui/", assets.Handler())

	return &Service{
		db:         db,
		srv:        srv,
		aggregator: am,
	}, nil
}
