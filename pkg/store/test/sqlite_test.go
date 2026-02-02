package test

import (
	"context"
	"os"
	"testing"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
	"github.com/stretchr/testify/assert"
)

func setupSQLiteTestCase(t *testing.T) (store.DB, func(t *testing.T)) {
	t.Log("setup SQLite test case")
	dbPath := "/tmp/feedpushr_test.db"

	// Remove existing test database
	os.Remove(dbPath)

	db, err := store.NewDB("sqlite://"+dbPath, model.Quota{})
	if err != nil {
		t.Fatalf("Unable to setup SQLite Database: %v", err)
	}

	return db, func(t *testing.T) {
		t.Log("teardown SQLite test case")
		db.Close()
		os.Remove(dbPath)
	}
}

func TestSQLiteFeedCRUD(t *testing.T) {
	db, teardown := setupSQLiteTestCase(t)
	defer teardown(t)

	ctx := context.Background()
	feed := &model.FeedDef{
		ID:     "test-sqlite",
		XMLURL: "http://example.com/feed.xml",
		Title:  "Test Feed",
		Tags:   []string{"test", "sqlite"},
	}

	// Test Save
	err := db.SaveFeed(ctx, feed)
	assert.Nil(t, err)

	// Test List
	page, err := db.ListFeeds(ctx, 1, 10)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Equal(t, 1, page.Page)
	assert.Equal(t, 10, page.Size)
	assert.Len(t, page.Feeds, 1, "unexpected number of feeds")
	assert.Equal(t, "test-sqlite", page.Feeds[0].ID, "unexpected feed ID")

	// Test Count
	total, err := db.CountFeeds(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 1, total, "unexpected number of feeds")

	// Test Get
	retrievedFeed, err := db.GetFeed(ctx, "test-sqlite")
	assert.Nil(t, err)
	assert.NotNil(t, retrievedFeed)
	assert.Equal(t, "test-sqlite", retrievedFeed.ID, "unexpected feed ID")
	assert.Equal(t, "Test Feed", retrievedFeed.Title, "unexpected feed title")
	assert.Equal(t, []string{"test", "sqlite"}, retrievedFeed.Tags, "unexpected feed tags")

	// Test Delete
	deletedFeed, err := db.DeleteFeed(ctx, "test-sqlite")
	assert.Nil(t, err)
	assert.NotNil(t, deletedFeed)
	assert.Equal(t, "test-sqlite", deletedFeed.ID)

	// Verify deletion
	_, err = db.GetFeed(ctx, "test-sqlite")
	assert.NotNil(t, err)
}

func TestSQLiteCacheCRUD(t *testing.T) {
	db, teardown := setupSQLiteTestCase(t)
	defer teardown(t)

	ctx := context.Background()
	// Test Store to cache
	err := db.StoreToCache(ctx, "test-key", &model.CacheItem{
		Value: "test-value",
	})
	assert.Nil(t, err)

	// Test Get from cache
	item, err := db.GetFromCache(ctx, "test-key")
	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, "test-value", item.Value)

	// Test Clear cache
	err = db.ClearCache(ctx)
	assert.Nil(t, err)

	// Verify cache is empty
	item, err = db.GetFromCache(ctx, "test-key")
	assert.Nil(t, err)
	assert.Nil(t, item)
}

func TestSQLiteFTSSearch(t *testing.T) {
	db, teardown := setupSQLiteTestCase(t)
	defer teardown(t)

	ctx := context.Background()
	// Create test feeds
	feeds := []*model.FeedDef{
		{
			ID:     "feed1",
			XMLURL: "http://example.com/golang.xml",
			Title:  "Golang News",
			Tags:   []string{"golang", "programming"},
		},
		{
			ID:     "feed2",
			XMLURL: "http://example.com/python.xml",
			Title:  "Python Daily",
			Tags:   []string{"python", "programming"},
		},
		{
			ID:     "feed3",
			XMLURL: "http://example.com/rust.xml",
			Title:  "Rust Weekly",
			Tags:   []string{"rust", "systems"},
		},
	}

	// Save feeds
	for _, feed := range feeds {
		err := db.SaveFeed(ctx, feed)
		assert.Nil(t, err)
	}

	// Build initial index
	err := db.BuildInitialIndex(ctx)
	assert.Nil(t, err)

	// Search for "golang"
	page, err := db.SearchFeeds(ctx, "golang", 1, 10)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Equal(t, 1, page.Total, "expected 1 result for 'golang'")
	assert.Len(t, page.Feeds, 1)
	assert.Equal(t, "Golang News", page.Feeds[0].Title)

	// Search for "programming"
	page, err = db.SearchFeeds(ctx, "programming", 1, 10)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Equal(t, 2, page.Total, "expected 2 results for 'programming'")
	assert.Len(t, page.Feeds, 2)

	// Search for "rust"
	page, err = db.SearchFeeds(ctx, "rust", 1, 10)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Equal(t, 1, page.Total, "expected 1 result for 'rust'")
	assert.Len(t, page.Feeds, 1)
	assert.Equal(t, "Rust Weekly", page.Feeds[0].Title)
}
