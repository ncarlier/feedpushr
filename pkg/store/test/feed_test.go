package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/v2/pkg/assert"
	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

func TestFeedCRUD(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	feed := &model.FeedDef{
		ID: "test",
	}
	err := db.SaveFeed(feed)
	assert.Nil(t, err, "should be nil")

	feeds, err := db.ListFeeds(1, 10)
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, feeds, "should not be nil")
	assert.Equal(t, 1, len(*feeds), "unexpected number of feeds")
	assert.Equal(t, "test", (*feeds)[0].ID, "unexpected feed ID")
	total, err := db.CountFeeds()
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, 1, total, "unexpected number of feeds")
	feed, err = db.GetFeed("test")
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, feed, "should not be nil")
	assert.Equal(t, "test", feed.ID, "unexpected feed ID")
	_, err = db.DeleteFeed("test")
	assert.Nil(t, err, "should be nil")
	_, err = db.GetFeed("test")
	assert.NotNil(t, err, "should not be nil")
	assert.Equal(t, common.ErrFeedNotFound.Error(), err.Error(), "unexpected error message")
}
