package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

func TestCacheCRUD(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	ctx := context.Background()
	item := &model.CacheItem{
		Value: "test",
	}
	err := db.StoreToCache(ctx, "test", item)
	assert.Nil(t, err)

	item, err = db.GetFromCache(ctx, "test")
	assert.Nil(t, err)
	assert.NotNil(t, item, "should not be nil")
	assert.Equal(t, "test", item.Value, "unexpected item value")
	err = db.ClearCache(ctx)
	assert.Nil(t, err)
	item, err = db.GetFromCache(ctx, "test")
	assert.Nil(t, err)
	assert.Nil(t, item)
}
