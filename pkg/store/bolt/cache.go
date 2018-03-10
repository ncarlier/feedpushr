package store

import (
	"encoding/json"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/ncarlier/feedpushr/pkg/model"
)

// CACHE_BUCKET bucket name
var CACHE_BUCKET = []byte("CACHE")

// GetFromCache returns a cached item.
func (store *BoltStore) GetFromCache(key string) (*model.CacheItem, error) {
	var result model.CacheItem
	err := store.get(CACHE_BUCKET, key, &result)
	if err != nil {
		if err == bolt.ErrInvalid {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

// StoreToCache stores a item into the cache.
func (store *BoltStore) StoreToCache(key string, item *model.CacheItem) error {
	return store.save(CACHE_BUCKET, key, &item)
}

// ClearCache removes all items from the cache.
func (store *BoltStore) ClearCache() error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(CACHE_BUCKET)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucket(CACHE_BUCKET)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// EvictFromCache manage the cache eviction.
func (store *BoltStore) EvictFromCache(before time.Time) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(CACHE_BUCKET)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var item model.CacheItem
			if err := json.Unmarshal(v, &item); err != nil {
				// Unable to decode? OK delete!
				b.Delete(k)
				continue
			}
			if item.Date.Before(before) {
				b.Delete(k)
			}
		}
		return nil
	})
	return err
}
