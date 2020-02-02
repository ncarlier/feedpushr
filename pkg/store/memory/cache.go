package store

import (
	"time"

	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

// GetFromCache returns a cached item.
func (store *InMemoryStore) GetFromCache(key string) (*model.CacheItem, error) {
	item, exists := store.cache[key]
	if !exists {
		return nil, nil
	}
	return &item, nil
}

// StoreToCache stores a item into the cache.
func (store *InMemoryStore) StoreToCache(key string, item *model.CacheItem) error {
	store.cacheLock.RLock()
	defer store.cacheLock.RUnlock()
	store.cache[key] = *item
	return nil
}

// ClearCache removes all items from the cache.
func (store *InMemoryStore) ClearCache() error {
	store.cacheLock.RLock()
	defer store.cacheLock.RUnlock()
	store.cache = make(map[string]model.CacheItem)
	return nil
}

// EvictFromCache manage the cache eviction.
func (store *InMemoryStore) EvictFromCache(before time.Time) error {
	store.cacheLock.RLock()
	defer store.cacheLock.RUnlock()
	for key, item := range store.cache {
		if item.Date.Before(before) {
			delete(store.cache, key)
		}
	}
	return nil
}
