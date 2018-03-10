package store

import (
	"fmt"
	"net/url"
	"time"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/cache"
	bolt "github.com/ncarlier/feedpushr/pkg/store/bolt"
	"github.com/rs/zerolog/log"
)

// DB is the data store
type DB interface {
	ListFeeds(page, limit int) (*app.FeedCollection, error)
	ExistsFeed(url string) bool
	GetFeed(id string) (*app.Feed, error)
	DeleteFeed(id string) (*app.Feed, error)
	SaveFeed(feed *app.Feed) error
	ForEachFeed(cb func(*app.Feed) error) error
	GetFromCache(key string) (*cache.Item, error)
	StoreToCache(key string, item *cache.Item) error
	ClearCache() error
	EvictFromCache(before time.Time) error
	Close() error
}

// Configure the data store regarding the datasource URI
func Configure(datasource string) (DB, error) {
	u, err := url.ParseRequestURI(datasource)
	if err != nil {
		return nil, fmt.Errorf("invalid datasource URL: %s", datasource)
	}
	datastore := u.Scheme
	var db DB

	switch datastore {
	case "boltdb":
		db, err = bolt.NewBoltStore(u)
		if err != nil {
			return nil, err
		}
		log.Info().Str("component", "store").Str("uri", u.String()).Msg("using BoltDB datastore")
	default:
		return nil, fmt.Errorf("unsuported datastore: %s", datastore)
	}
	return db, nil
}
