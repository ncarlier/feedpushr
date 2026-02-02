package store

import (
	"context"
	"time"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// CacheRepository interface to manage cache
type CacheRepository interface {
	GetFromCache(ctx context.Context, key string) (*model.CacheItem, error)
	StoreToCache(ctx context.Context, key string, item *model.CacheItem) error
	ClearCache(ctx context.Context) error
	EvictFromCache(ctx context.Context, before time.Time) error
}
