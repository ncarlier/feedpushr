package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/model"
)

func TestCacheCRUD(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	item := &model.CacheItem{
		Value: "test",
	}
	err := db.StoreToCache("test", item)
	assert.Nil(t, err, "should be nil")

	item, err = db.GetFromCache("test")
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, item, "should not be nil")
	assert.Equal(t, "test", item.Value, "unexpected item value")
	err = db.ClearCache()
	assert.Nil(t, err, "should be nil")
	item, err = db.GetFromCache("test")
	assert.Nil(t, err, "should be nil")
	assert.Nil(t, item, "should be nil")
}
