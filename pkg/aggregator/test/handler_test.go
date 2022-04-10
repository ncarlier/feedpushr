package aggregator_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v3/pkg/builder"
)

const feedTitle = "Le Monde.fr - Actualités et Infos en France et dans le monde"

func TestNewFeedHandler(t *testing.T) {
	url := "https://www.lemonde.fr/rss/une.xml"
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
	article := items[0]
	assert.Equal(t, feedTitle, article.FeedTitle)
	assert.NotEmpty(t, article.Title)
	status, items = handler.Refresh()
	assert.NotNil(t, status)
	assert.Empty(t, status.ErrorMsg)
	assert.Empty(t, items)
}
