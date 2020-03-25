package store

import (
	"sync"

	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

// InMemoryStore is a data store backed by InMemoryDB
type InMemoryStore struct {
	cache       map[string]model.CacheItem
	cacheLock   sync.RWMutex
	feeds       map[string]model.FeedDef
	feedsLock   sync.RWMutex
	outputs     map[string]model.OutputDef
	outputsLock sync.RWMutex
}

// NewInMemoryStore creates a data store backed by InMemoryDB
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		cache:       make(map[string]model.CacheItem),
		cacheLock:   sync.RWMutex{},
		feeds:       make(map[string]model.FeedDef),
		feedsLock:   sync.RWMutex{},
		outputs:     make(map[string]model.OutputDef),
		outputsLock: sync.RWMutex{},
	}
}

// Close the DB
func (store *InMemoryStore) Close() error {
	return nil
}
