package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/ncarlier/feedpushr/v3/pkg/aggregator"
	pkgapi "github.com/ncarlier/feedpushr/v3/pkg/api"
	"github.com/ncarlier/feedpushr/v3/pkg/auth"
	"github.com/ncarlier/feedpushr/v3/pkg/cache"
	"github.com/ncarlier/feedpushr/v3/pkg/config"
	"github.com/ncarlier/feedpushr/v3/pkg/controller"
	"github.com/ncarlier/feedpushr/v3/pkg/explore"
	"github.com/ncarlier/feedpushr/v3/pkg/filter"
	"github.com/ncarlier/feedpushr/v3/pkg/handler"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/output"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
	"github.com/rs/zerolog/log"
)

// Server instance
type Server struct {
	conf       config.Config
	db         store.DB
	httpServer *http.Server
	aggregator *aggregator.Manager
	outputs    *output.Manager
	cache      *cache.Manager
	listener   net.Listener
	agent      *api.Agent
}

// ListenAndServe starts server
func (s *Server) ListenAndServe(listenAddr string) error {
	log.Debug().Msg("loading output manager...")
	if err := loadOutputs(s.db, s.outputs); err != nil {
		return err
	}
	log.Debug().Msg("loading feed aggregators...")
	if err := loadFeedAggregators(s.db, s.aggregator, s.conf.FanOutDelay); err != nil {
		return err
	}
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	s.listener = listener
	if err := s.register(); err != nil {
		log.Debug().Err(err).Msg("unable to register service")
	}

	log.Debug().Msg("starting HTTP server...")
	if err := s.httpServer.Serve(s.listener); err != nil && err != http.ErrServerClosed {
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
	return s.httpServer.Shutdown(ctx)
}

// NewServer creates new server instance
func NewServer(db store.DB, conf config.Config) (*Server, error) {
	// Load plugins
	if err := loadPlugins(conf); err != nil {
		return nil, err
	}

	// Clear configuration if asked
	if conf.ClearConfig {
		if err := db.ClearOutputs(context.Background()); err != nil {
			log.Error().Err(err).Msg("unable to clear outputs")
			return nil, err
		}
	}

	// Init feed explorer
	explorer, err := explore.NewExplorer(conf.ExploreProvider)
	if err != nil {
		return nil, err
	}

	// Create empty chain filter (for filter controller)
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

	// Create authenticator
	issuer := ""
	authenticator, err := auth.NewAuthenticator(conf.Authn, conf.AuthorizedUsername)
	if err != nil {
		log.Info().Err(err).Str("authn", conf.Authn).Msg("unable to load authenticator")
		authenticator = nil
	} else if authenticator != nil {
		issuer = authenticator.Issuer()
		log.Info().Str("authn", conf.Authn).Msg("using authenticator")
	}

	// Create router
	router := pkgapi.NewRouter()

	// Create handlers
	h := handler.New(db, am, om, explorer, cf, issuer, conf.ClientID)

	// Register all API routes
	h.RegisterRoutes(router)

	// Mount custom handlers (UI and redirects)
	router.Handle("GET", "/ui/{path...}", controller.NewUIHandler())
	router.HandleFunc("GET", "/", controller.NewRedirectHandler(conf.PublicURL+"/ui/"))

	// Build middleware chain
	authMiddleware := auth.GetAuthMiddleware(authenticator, "/v2/", "/v2/healthz", "/v2/pshb")

	middlewares := pkgapi.Chain(
		pkgapi.RequestIDMiddleware,
		pkgapi.LoggingMiddleware,
		pkgapi.RecoverMiddleware,
		pkgapi.CORSMiddleware,
		authMiddleware,
	)

	// Create HTTP server
	httpServer := &http.Server{
		Handler:      middlewares(router),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{
		db:         db,
		httpServer: httpServer,
		conf:       conf,
		aggregator: am,
		outputs:    om,
		cache:      cm,
	}, nil
}
