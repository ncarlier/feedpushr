package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ncarlier/feedpushr/v3/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v3/pkg/api"
	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/feed"
	"github.com/ncarlier/feedpushr/v3/pkg/helper"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// CreateFeedRequest represents a create feed request
type CreateFeedRequest struct {
	URL    string   `json:"url"`
	Title  *string  `json:"title,omitempty"`
	Tags   []string `json:"tags,omitempty"`
	Enable *bool    `json:"enable,omitempty"`
}

// UpdateFeedRequest represents an update feed request
type UpdateFeedRequest struct {
	Title *string  `json:"title,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

// CreateFeed handles POST /v2/feeds
func (h *Handlers) CreateFeed(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Decode request payload
	var payload CreateFeedRequest
	if err := req.Decode(&payload); err != nil {
		res.BadRequest(err)
		return
	}

	// Create feed definition
	_tags := tagsToString(payload.Tags)
	def, err := feed.NewFeed(payload.URL, _tags)
	if err != nil {
		res.BadRequest(err)
		return
	}

	// Set optional title
	if !helper.IsEmptyString(payload.Title) {
		def.Title = *payload.Title
	}

	// Set initial status
	status := aggregator.StoppedStatus.String()
	if payload.Enable != nil && *payload.Enable {
		status = aggregator.RunningStatus.String()
	}
	def.Status = &status

	// Save feed definition
	if err := h.db.SaveFeed(r.Context(), def); err != nil {
		res.InternalError(err)
		return
	}

	// Start feed aggregation if enabled
	if def.Status != nil && *def.Status == aggregator.RunningStatus.String() {
		fa := h.aggregator.RegisterFeedAggregator(def, 0)
		fa.Start()
		h.log.Info().Str("id", def.ID).Msg("feed created and registered")
	} else {
		h.log.Info().Str("id", def.ID).Msg("feed created")
	}

	res.Created(def)
}

// UpdateFeed handles PUT /v2/feeds/{id}
func (h *Handlers) UpdateFeed(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Retrieve feed ID from path
	id := req.PathParam("id")

	// Decode request payload
	var payload UpdateFeedRequest
	if err := req.Decode(&payload); err != nil {
		res.BadRequest(err)
		return
	}

	// Retrieve existing feed definition
	def, err := h.db.GetFeed(r.Context(), id)
	if err != nil {
		if err == common.ErrFeedNotFound {
			res.NotFound()
			return
		}
		res.InternalError(err)
		return
	}

	// If no update fields provided, return existing feed
	if payload.Tags == nil && helper.IsEmptyString(payload.Title) {
		res.OK(def)
		return
	}

	// Update feed definition
	if !helper.IsEmptyString(payload.Title) {
		def.Title = *payload.Title
	}
	tags := tagsToString(payload.Tags)
	def.Tags = feed.GetFeedTags(tags)
	def.Mdate = time.Now()

	// Save updated feed definition
	if err := h.db.SaveFeed(r.Context(), def); err != nil {
		res.InternalError(err)
		return
	}

	// Restart feed aggregation if running
	fa := h.aggregator.GetFeedAggregator(def.ID)
	if fa != nil {
		h.aggregator.UnRegisterFeedAggregator(def.ID)
		if def.Status != nil && *def.Status == aggregator.RunningStatus.String() {
			h.aggregator.RegisterFeedAggregator(def, 0)
			h.log.Info().Str("id", def.ID).Msg("feed updated and aggregation restarted")
		} else {
			h.log.Info().Str("id", def.ID).Msg("feed updated and aggregation stopped")
		}
	} else {
		h.log.Info().Str("id", def.ID).Msg("feed updated")
	}

	res.OK(def)
}

// DeleteFeed handles DELETE /v2/feeds/{id}
func (h *Handlers) DeleteFeed(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Retrieve feed ID from path
	id := req.PathParam("id")

	// Unregister feed aggregator and delete feed definition
	h.aggregator.UnRegisterFeedAggregator(id)
	_, err := h.db.DeleteFeed(r.Context(), id)
	if err != nil {
		if err == common.ErrFeedNotFound {
			res.NotFound()
			return
		}
		res.InternalError(err)
		return
	}

	h.log.Info().Str("id", id).Msg("feed removed and aggregation stopped")
	res.NoContent()
}

// ListFeeds handles GET /v2/feeds
func (h *Handlers) ListFeeds(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Get pagination and query parameters
	page, size := req.Page()
	q := req.QueryParam("q")

	var feedPage *model.FeedDefPage
	var err error

	// Retrieve feed definitions from database
	if q != "" {
		// Search feeds
		feedPage, err = h.db.SearchFeeds(r.Context(), q, page, size)
	} else {
		// List feeds
		feedPage, err = h.db.ListFeeds(r.Context(), page, size)
	}

	if err != nil {
		res.InternalError(err)
		return
	}

	// Convert to response format
	response := convertFeedPageToResponse(feedPage, h.aggregator)
	res.OK(response)
}

// GetFeed handles GET /v2/feeds/{id}
func (h *Handlers) GetFeed(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Retrieve feed ID from path
	id := req.PathParam("id")

	// Retrieve existing feed definition
	def, err := h.db.GetFeed(r.Context(), id)
	if err != nil {
		if err == common.ErrFeedNotFound {
			res.NotFound()
			return
		}
		res.InternalError(err)
		return
	}

	// Get feed aggregator if registered
	fa := h.aggregator.GetFeedAggregator(def.ID)
	if fa != nil {
		res.OK(fa.GetFeedWithAggregationStatus())
		return
	}

	res.OK(def)
}

// StartFeed handles POST /v2/feeds/{id}/start
func (h *Handlers) StartFeed(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Retrieve feed ID from path
	id := req.PathParam("id")

	// Retrieve existing feed definition
	def, err := h.db.GetFeed(r.Context(), id)
	if err != nil {
		if err == common.ErrFeedNotFound {
			res.NotFound()
			return
		}
		res.InternalError(err)
		return
	}

	// Start feed aggregation
	fa := h.aggregator.GetFeedAggregator(def.ID)
	if fa == nil {
		fa = h.aggregator.RegisterFeedAggregator(def, 0)
	}
	fa.StartWithDelay(0)

	// Update feed status
	status := aggregator.RunningStatus.String()
	def.Status = &status
	if err := h.db.SaveFeed(r.Context(), def); err != nil {
		res.InternalError(err)
		return
	}

	h.log.Info().Str("id", def.ID).Msg("feed aggregation started")
	res.Accepted()
}

// StopFeed handles POST /v2/feeds/{id}/stop
func (h *Handlers) StopFeed(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Retrieve feed ID from path
	id := req.PathParam("id")

	// Retrieve existing feed definition
	def, err := h.db.GetFeed(r.Context(), id)
	if err != nil {
		if err == common.ErrFeedNotFound {
			res.NotFound()
			return
		}
		res.InternalError(err)
		return
	}

	// Stop feed aggregation
	fa := h.aggregator.GetFeedAggregator(def.ID)
	if fa == nil {
		res.InternalError(fmt.Errorf("feed aggregator not registered"))
		return
	}
	fa.Stop()

	// Update feed status
	status := aggregator.StoppedStatus.String()
	def.Status = &status
	if err := h.db.SaveFeed(r.Context(), def); err != nil {
		res.InternalError(err)
		return
	}

	h.log.Info().Str("id", def.ID).Msg("feed aggregation stopped")
	res.Accepted()
}

// convertFeedPageToResponse converts feed page to response format
func convertFeedPageToResponse(page *model.FeedDefPage, am *aggregator.Manager) interface{} {
	response := api.FeedsPageResponse{
		Current: page.Page,
		Size:    page.Size,
		Total:   page.Total,
		Data:    make([]api.FeedResponse, 0, len(page.Feeds)),
	}

	for _, f := range page.Feeds {
		agg := am.GetFeedAggregator(f.ID)
		if agg == nil {
			response.Data = append(response.Data, api.FeedResponse{FeedDef: f})
		} else {
			response.Data = append(response.Data, *agg.GetFeedWithAggregationStatus())
		}
	}

	return response
}

// tagsToString converts a string slice to a pointer to comma-separated string
func tagsToString(tags []string) *string {
	if len(tags) == 0 {
		return nil
	}
	result := ""
	for i, tag := range tags {
		if i > 0 {
			result += ","
		}
		result += tag
	}
	return &result
}
