package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// GetFromCache returns a cached item.
func (store *SQLiteStore) GetFromCache(ctx context.Context, key string) (*model.CacheItem, error) {
	var item model.CacheItem

	err := store.db.QueryRowContext(ctx, `
		SELECT value, date FROM cache WHERE key = ?
	`, key).Scan(&item.Value, &item.Date)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}

// StoreToCache stores a item into the cache.
func (store *SQLiteStore) StoreToCache(ctx context.Context, key string, item *model.CacheItem) error {
	_, err := store.db.ExecContext(ctx, `
		INSERT INTO cache (key, value, date)
		VALUES (?, ?, ?)
		ON CONFLICT(key) DO UPDATE SET
			value = excluded.value,
			date = excluded.date
	`, key, item.Value, item.Date)

	return err
}

// ClearCache removes all items from the cache.
func (store *SQLiteStore) ClearCache(ctx context.Context) error {
	_, err := store.db.ExecContext(ctx, "DELETE FROM cache")
	return err
}

// EvictFromCache manage the cache eviction.
func (store *SQLiteStore) EvictFromCache(ctx context.Context, before time.Time) error {
	_, err := store.db.ExecContext(ctx, "DELETE FROM cache WHERE date < ?", before)
	return err
}
