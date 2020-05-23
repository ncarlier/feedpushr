package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

func TestFeedCRUD(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	feed := &model.FeedDef{
		ID: "test",
	}
	err := db.SaveFeed(feed)
	assert.Nil(t, err)

	page, err := db.ListFeeds(1, 10)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Equal(t, 1, page.Page)
	assert.Equal(t, 10, page.Size)
	assert.Len(t, page.Feeds, 1, "unexpected number of feeds")
	assert.Equal(t, "test", page.Feeds[0].ID, "unexpected feed ID")
	assert.Equal(t, 1, page.Total, "unexpected number of total feeds")
	total, err := db.CountFeeds()
	assert.Nil(t, err)
	assert.Equal(t, 1, total, "unexpected number of feeds")
	feed, err = db.GetFeed("test")
	assert.Nil(t, err)
	assert.NotNil(t, feed)
	assert.Equal(t, "test", feed.ID, "unexpected feed ID")
	_, err = db.DeleteFeed("test")
	assert.Nil(t, err)
	_, err = db.GetFeed("test")
	assert.NotNil(t, err)
	assert.Equal(t, common.ErrFeedNotFound.Error(), err.Error(), "unexpected error message")
}
