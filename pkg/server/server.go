package server

import (
	"context"
	"net"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	consul "github.com/hashicorp/consul/api"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
	"github.com/ncarlier/feedpushr/v3/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v3/pkg/assets"
	"github.com/ncarlier/feedpushr/v3/pkg/auth"
	"github.com/ncarlier/feedpushr/v3/pkg/cache"
	"github.com/ncarlier/feedpushr/v3/pkg/config"
	"github.com/ncarlier/feedpushr/v3/pkg/controller"
	"github.com/ncarlier/feedpushr/v3/pkg/explore"
	"github.com/ncarlier/feedpushr/v3/pkg/filter"
	"github.com/ncarlier/feedpushr/v3/pkg/logging"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/output"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
	"github.com/rs/zerolog/log"
)

// Server instance
type Server struct {
	conf       config.Config
	db         store.DB
	srv        *goa.Service
	aggregator *aggregator.Manager
	outputs    *output.Manager
	cache      *cache.Manager
	listener   net.Listener
	agent      *consul.Agent
}

// ListenAndServe starts server
func (s *Server) ListenAndServe(ListenAddr string) error {
	log.Debug().Msg("loading output manager...")
	if err := loadOutputs(s.db, s.outputs); err != nil {
		return err
	}
	log.Debug().Msg("loading feed aggregators...")
	if err := loadFeedAggregators(s.db, s.aggregator, s.conf.FanOutDelay); err != nil {
		return err
	}
	listener, err := net.Listen("tcp", ListenAddr)
	if err != nil {
		return err
	}
	s.listener = listener
	if err := s.register(); err != nil {
		log.Debug().Err(err).Msg("unable to register service")
	}

	log.Debug().Msg("starting HTTP server...")
	if err := s.srv.Serve(s.listener); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Shutdown server and managed service
func (s *Server) Shutdown(ctx context.Context) error {
	s.cache.Shutdown()
	s.aggregator.Shutdown()
	s.outputs.Shutdown()
	s.deregister()
	s.srv.CancelAll()
	s.srv.Server.SetKeepAlivesEnabled(false)
	return s.srv.Server.Shutdown(ctx)
}

// NewServer creates new server instance
func NewServer(db store.DB, conf config.Config) (*Server, error) {
	// Load plugins
	if err := loadPlugins(conf); err != nil {
		return nil, err
	}

	// Clear configuration if asked
	if conf.ClearConfig {
		if err := db.ClearOutputs(); err != nil {
			log.Error().Err(err).Msg("unable to clear outputs")
			return nil, err
		}
	}

	// Init feed explorer
	explorer, err := explore.NewExplorer(conf.ExploreProvider)
	if err != nil {
		return nil, err
	}

	// Creat empty chain filter (for filter controller)
	cf, err := filter.NewChainFilter(model.FilterDefCollection{})
	if err != nil {
		return nil, err
	}

	// Init cache manager
	cm, err := cache.NewCacheManager(db, conf)
	if err != nil {
		return nil, err
	}

	// Init output manager
	om, err := output.NewOutputManager(cm)
	if err != nil {
		log.Error().Err(err).Msg("unable to init output manager")
		return nil, err
	}

	// Init aggregator daemon
	var callbackURL string
	if conf.PublicURL != "" {
		callbackURL = conf.PublicURL + "/v2/pshb"
	}
	am := aggregator.NewAggregatorManager(om, conf.Delay, conf.Timeout, callbackURL)

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
	htpasswd, err := auth.NewHtpasswdFromFile(conf.PasswdFile)
	if err != nil {
		log.Debug().Err(err).Msg("unable to load htpasswd file: authentication deactivated")
	} else {
		srv.Use(auth.NewMiddleware(htpasswd, "/pshb"))
	}

	// Mount "index" controller
	app.MountIndexController(srv, controller.NewIndexController(srv))
	// Mount "feed" controller
	app.MountFeedController(srv, controller.NewFeedController(srv, db, am))
	// Mount "filter" controller
	app.MountFilterController(srv, controller.NewFilterController(srv, cf))
	// Mount "output" controller
	app.MountOutputController(srv, controller.NewOutputController(srv, db, om))
	// Mount "health" controller
	app.MountHealthController(srv, controller.NewHealthController(srv))
	// Mount "swagger" controller
	app.MountSwaggerController(srv, controller.NewSwaggerController(srv))
	// Mount "opml" controller
	app.MountOpmlController(srv, controller.NewOpmlController(srv, db))
	// Mount "explore" controller
	app.MountExploreController(srv, controller.NewExploreController(srv, explorer))
	// Mount "vars" controller
	app.MountVarsController(srv, controller.NewVarsController(srv))
	// Mount "pshb" controller (only if public URL is configured)
	if conf.PublicURL != "" {
		app.MountPshbController(srv, controller.NewPshbController(srv, db, am, om))
	}
	// Mount custom handlers (aka: not generated)...
	srv.Mux.Handle("GET", "/ui/*asset", assets.Handler())
	srv.Mux.Handle("GET", "/ui/", assets.Handler())
	srv.Mux.Handle("GET", "/", controller.Redirect(conf.PublicURL+"/ui/"))

	return &Server{
		db:         db,
		srv:        srv,
		conf:       conf,
		aggregator: am,
		outputs:    om,
		cache:      cm,
	}, nil
}
