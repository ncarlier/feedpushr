package store

import (
	"fmt"
	"net/url"

	bolt "github.com/ncarlier/feedpushr/v3/pkg/store/bolt"
	memory "github.com/ncarlier/feedpushr/v3/pkg/store/memory"
	"github.com/rs/zerolog/log"
)

// DB is the interface with the database
type DB interface {
	FeedRepository
	OutputRepository
	CacheRepository
	Close() error
}

// NewDB creates new database access regarding the datasource URI
func NewDB(datasource string) (DB, error) {
	u, err := url.ParseRequestURI(datasource)
	if err != nil {
		return nil, fmt.Errorf("invalid datasource URL: %s", datasource)
	}
	provider := u.Scheme
	var db DB

	switch provider {
	case "memory":
		db = memory.NewInMemoryStore()
		log.Info().Str("component", "db").Str("uri", u.String()).Msg("using in memory database")
	case "boltdb":
		db, err = bolt.NewBoltStore(u)
		if err != nil {
			return nil, err
		}
		log.Info().Str("component", "db").Str("uri", u.String()).Msg("using BoltDB")
	default:
		return nil, fmt.Errorf("unsupported database provider: %s", provider)
	}
	return db, nil
}
