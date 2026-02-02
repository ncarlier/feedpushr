package store

import (
	"context"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// FeedRepository interface to manage feeds
type FeedRepository interface {
	ListFeeds(ctx context.Context, page, size int) (*model.FeedDefPage, error)
	CountFeeds(ctx context.Context) (int, error)
	ExistsFeed(ctx context.Context, url string) bool
	GetFeed(ctx context.Context, id string) (*model.FeedDef, error)
	DeleteFeed(ctx context.Context, id string) (*model.FeedDef, error)
	SaveFeed(ctx context.Context, feed *model.FeedDef) error
	ForEachFeed(ctx context.Context, cb func(*model.FeedDef) error) error
}
