package api

import (
	"time"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// FeedResponse represents a feed in API responses
type FeedResponse struct {
	model.FeedDef
	NbProcessedItems *int       `json:"nbProcessedItems,omitempty"`
	NbErrors         *int       `json:"nbErrors,omitempty"`
	LastErrorMsg     *string    `json:"lastErrorMsg,omitempty"`
	NextCheck        *time.Time `json:"nextCheck,omitempty"`
}

// FeedsPageResponse represents a paginated list of feeds
type FeedsPageResponse struct {
	Size    int            `json:"size"`
	Current int            `json:"current"`
	Total   int            `json:"total"`
	Data    []FeedResponse `json:"data"`
}

// HALLink represents a HAL hypermedia link
type HALLink struct {
	Href string `json:"href"`
}

// Info represents API information
type Info struct {
	Name     string              `json:"name"`
	Desc     string              `json:"desc"`
	Version  string              `json:"version"`
	ClientID string              `json:"client_id,omitempty"`
	Links    map[string]*HALLink `json:"_links,omitempty"`
}
