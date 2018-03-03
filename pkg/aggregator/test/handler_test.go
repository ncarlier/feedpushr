package aggregator_test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/aggregator"
	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/builder"
)

func TestNewFeedHandler(t *testing.T) {
	url := "https://keeper.nunux.org/index.xml"
	feed, err := builder.NewFeed(url)
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, feed, "feed shouldn't be nil")

	handler := aggregator.NewFeedHandler(feed)
	status, items := handler.Refresh()
	assert.NotNil(t, status, "items shouldn't be nil")
	assert.NotNil(t, items, "items feed shouldn't be nil")
	assert.Equal(t, "", status.ErrorMsg, "status error message should be empty")
	assert.Equal(t, 0, status.ErrorCount, "status error count should be 0")
	assert.True(t, len(items) > 0, "items shouldn't be empty")
	status, items = handler.Refresh()
	assert.NotNil(t, status, "items shouldn't be nil")
	assert.NotNil(t, items, "items feed shouldn't be nil")
	assert.Equal(t, "", status.ErrorMsg, "status error message should be empty")
	assert.Equal(t, 0, len(items), "items should be empty")
}
