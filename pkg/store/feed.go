package store

import (
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// FeedRepository interface to manage feeds
type FeedRepository interface {
	ListFeeds(page, limit int) (*model.FeedDefCollection, error)
	CountFeeds() (int, error)
	ExistsFeed(url string) bool
	GetFeed(id string) (*model.FeedDef, error)
	DeleteFeed(id string) (*model.FeedDef, error)
	SaveFeed(feed *model.FeedDef) error
	ForEachFeed(cb func(*model.FeedDef) error) error
}
