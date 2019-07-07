package store

import (
	"sync"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/model"
)

// InMemoryStore is a data store backed by InMemoryDB
type InMemoryStore struct {
	cache       map[string]model.CacheItem
	cacheLock   sync.RWMutex
	feeds       map[string]app.Feed
	feedsLock   sync.RWMutex
	filters     map[int]app.Filter
	filtersLock sync.RWMutex
	outputs     map[int]app.Output
	outputsLock sync.RWMutex
}

// NewInMemoryStore creates a data store backed by InMemoryDB
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		cache:       make(map[string]model.CacheItem),
		cacheLock:   sync.RWMutex{},
		feeds:       make(map[string]app.Feed),
		feedsLock:   sync.RWMutex{},
		filters:     make(map[int]app.Filter),
		filtersLock: sync.RWMutex{},
		outputs:     make(map[int]app.Output),
		outputsLock: sync.RWMutex{},
	}
}

// Close the DB
func (store *InMemoryStore) Close() error {
	return nil
}
