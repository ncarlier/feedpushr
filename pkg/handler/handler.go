package handler

import (
	"net/http"

	"github.com/ncarlier/feedpushr/v3/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v3/pkg/api"
	"github.com/ncarlier/feedpushr/v3/pkg/explore"
	"github.com/ncarlier/feedpushr/v3/pkg/filter"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/output"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Handlers contains all API handlers
type Handlers struct {
	db          store.DB
	aggregator  *aggregator.Manager
	outputs     *output.Manager
	explorer    explore.Explorer
	filterSpecs []model.Spec // Filter specifications
	issuer      string
	clientID    string
	log         zerolog.Logger
}

// New creates new handlers instance
func New(
	db store.DB,
	am *aggregator.Manager,
	om *output.Manager,
	explorer explore.Explorer,
	chainFilter *filter.Chain,
	issuer string,
	clientID string,
) *Handlers {
	return &Handlers{
		db:          db,
		aggregator:  am,
		outputs:     om,
		explorer:    explorer,
		filterSpecs: chainFilter.GetAvailableFilters(),
		issuer:      issuer,
		clientID:    clientID,
		log:         log.With().Str("component", "handlers").Logger(),
	}
}

// RegisterRoutes registers all API routes
func (h *Handlers) RegisterRoutes(r *api.Router) {
	// API root
	r.GET("/v2/", h.GetIndex)

	// Health check
	r.GET("/v2/healthz", h.GetHealth)

	// Feed endpoints
	r.GET("/v2/feeds", h.ListFeeds)
	r.POST("/v2/feeds", h.CreateFeed)
	r.GET("/v2/feeds/{id}", h.GetFeed)
	r.PUT("/v2/feeds/{id}", h.UpdateFeed)
	r.DELETE("/v2/feeds/{id}", h.DeleteFeed)
	r.POST("/v2/feeds/{id}/start", h.StartFeed)
	r.POST("/v2/feeds/{id}/stop", h.StopFeed)

	// Output endpoints
	r.GET("/v2/outputs", h.ListOutputs)
	r.POST("/v2/outputs", h.CreateOutput)
	r.GET("/v2/outputs/{id}", h.GetOutput)
	r.PUT("/v2/outputs/{id}", h.UpdateOutput)
	r.DELETE("/v2/outputs/{id}", h.DeleteOutput)
	r.GET("/v2/outputs/_specs", h.GetOutputSpecs)

	// Filter endpoints
	r.POST("/v2/outputs/{id}/filters", h.CreateFilter)
	r.PUT("/v2/outputs/{id}/filters/{fid}", h.UpdateFilter)
	r.DELETE("/v2/outputs/{id}/filters/{fid}", h.DeleteFilter)
	r.GET("/v2/filters/_specs", h.GetFilterSpecs)

	// OPML endpoints
	r.GET("/v2/opml", h.GetOPML)
	r.POST("/v2/opml", h.UploadOPML)
	r.GET("/v2/opml/status/{id}", h.GetOPMLStatus)

	// Explore endpoints
	r.GET("/v2/explore", h.ExploreFeeds)

	// Vars endpoints
	r.GET("/v2/vars", h.GetVars)

	// OpenAPI endpoints
	r.GET("/v2/openapi.json", h.GetOpenAPI)

	// PubSubHubbub endpoints
	r.GET("/v2/pshb", h.PshbSubscribe)
	r.POST("/v2/pshb", h.PshbPublish)
}

// Placeholder handlers - will be implemented in separate files
func (h *Handlers) GetIndex(w http.ResponseWriter, r *http.Request) {
	res := api.NewResponse(w)
	res.JSON(http.StatusOK, map[string]interface{}{
		"name":      "feedpushr",
		"desc":      "Feed aggregator daemon with sugar on top",
		"version":   "3.0.0",
		"client_id": h.clientID,
		"_links": map[string]interface{}{
			"documentation": map[string]string{"href": "https://github.com/ncarlier/feedpushr"},
		},
	})
}

func (h *Handlers) GetHealth(w http.ResponseWriter, r *http.Request) {
	res := api.NewResponse(w)
	res.Raw("text/plain", []byte("ok"))
}
