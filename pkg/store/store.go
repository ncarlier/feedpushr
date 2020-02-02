package store

import (
	"fmt"
	"net/url"

	bolt "github.com/ncarlier/feedpushr/v2/pkg/store/bolt"
	memory "github.com/ncarlier/feedpushr/v2/pkg/store/memory"
	"github.com/rs/zerolog/log"
)

// DB is the data store
type DB interface {
	FeedRepository
	FilterRepository
	OutputRepository
	CacheRepository
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
	case "memory":
		db = memory.NewInMemoryStore()
		log.Info().Str("component", "store").Str("uri", u.String()).Msg("using in memory datastore")
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
