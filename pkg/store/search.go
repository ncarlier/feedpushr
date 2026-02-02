package store

import (
	"context"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// SearchRepository interface to manage search index
type SearchRepository interface {
	BuildInitialIndex(ctx context.Context) error
	SearchFeeds(ctx context.Context, query string, page, size int) (*model.FeedDefPage, error)
}
