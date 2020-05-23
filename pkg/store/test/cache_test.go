package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

func TestCacheCRUD(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	item := &model.CacheItem{
		Value: "test",
	}
	err := db.StoreToCache("test", item)
	assert.Nil(t, err)

	item, err = db.GetFromCache("test")
	assert.Nil(t, err)
	assert.NotNil(t, item, "should not be nil")
	assert.Equal(t, "test", item.Value, "unexpected item value")
	err = db.ClearCache()
	assert.Nil(t, err)
	item, err = db.GetFromCache("test")
	assert.Nil(t, err)
	assert.Nil(t, item)
}
