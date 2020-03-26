package model

import "time"

// FeedDef object structure
type FeedDef struct {
	// ID of feed (MD5 of the xmlUrl)
	ID string `json:"id"`
	// URL of the XML feed
	XMLURL string `json:"xmlUrl"`
	// URL of the feed website
	HTMLURL *string `json:"htmlUrl,omitempty"`
	// URL of the PubSubHubbud hub
	HubURL *string `json:"hubUrl,omitempty"`
	// Title of the Feed
	Title string `json:"title"`
	// Aggregation status
	Status *string `json:"status,omitempty"`
	// List of tags
	Tags []string `json:"tags,omitempty"`
	// Date of creation
	Cdate time.Time `json:"cdate"`
	// Date of modification
	Mdate time.Time `json:"mdate"`
}

// FeedDefCollection is a list of feed definition
type FeedDefCollection []FeedDef
