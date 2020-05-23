package aggregator_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v3/pkg/builder"
)

func TestNewFeedHandler(t *testing.T) {
	url := "https://keeper.nunux.org/index.xml"
	feed, err := builder.NewFeed(url, nil)
	assert.Nil(t, err)
	assert.NotNil(t, feed)

	timeout := time.Duration(5) * time.Second
	handler := aggregator.NewFeedHandler(feed, timeout)
	status, items := handler.Refresh()
	assert.NotNil(t, status)
	assert.NotNil(t, items)
	assert.Empty(t, status.ErrorMsg)
	assert.Equal(t, 0, status.ErrorCount)
	assert.NotEmpty(t, items)
	status, items = handler.Refresh()
	assert.NotNil(t, status)
	assert.Empty(t, status.ErrorMsg)
	assert.Empty(t, items)
}
