package store

import (
	"time"

	"github.com/ncarlier/feedpushr/pkg/model"
)

// CacheRepository interface to manage cache
type CacheRepository interface {
	GetFromCache(key string) (*model.CacheItem, error)
	StoreToCache(key string, item *model.CacheItem) error
	ClearCache() error
	EvictFromCache(before time.Time) error
}
