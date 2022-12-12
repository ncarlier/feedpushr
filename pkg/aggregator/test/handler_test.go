package aggregator_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ncarlier/feedpushr/v3/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v3/pkg/builder"
)

var tests = []struct {
	URL   string
	title string
}{
	{
		URL:   "https://www.lemonde.fr/rss/une.xml",
		title: "Le Monde.fr - Actualités et Infos en France et dans le monde",
	},
	{
		URL:   "https://www.reddit.com/r/selfhosted.rss",
		title: "Self-Hosted Alternatives to Popular Services",
	},
}

func testFeedHandler(t *testing.T, url, title string) *aggregator.FeedHandler {
	feed, err := builder.NewFeed(url, nil)
	require.Nil(t, err)
	require.NotNil(t, feed)
	timeout := time.Duration(5) * time.Second
	handler := aggregator.NewFeedHandler(feed, timeout)
	status, items := handler.Refresh()
	require.NotNil(t, status)
	require.NotNil(t, items)
	require.Empty(t, status.ErrorMsg)
	require.Equal(t, 0, status.ErrorCount)
	require.NotEmpty(t, items)
	article := items[0]
	require.Equal(t, title, article.FeedTitle)
	require.NotEmpty(t, article.Title)
	return handler
}

func TestFeedHandler(t *testing.T) {
	for _, tt := range tests {
		testFeedHandler(t, tt.URL, tt.title)
	}
}

func TestFeedHandlerWithCacheHeaderSupport(t *testing.T) {
	tt := tests[0]
	handler := testFeedHandler(t, tt.URL, tt.title)
	status, items := handler.Refresh()
	require.NotNil(t, status)
	require.Empty(t, status.ErrorMsg)
	require.Empty(t, items)
}
