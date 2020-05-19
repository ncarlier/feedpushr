package store

import "github.com/ncarlier/feedpushr/v3/pkg/model"

// SearchRepository interface to manage search index
type SearchRepository interface {
	BuildInitialIndex() error
	SearchFeeds(query string, page, size int) (*model.FeedDefPage, error)
}
