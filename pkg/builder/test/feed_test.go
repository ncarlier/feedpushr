package builder_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ncarlier/feedpushr/v3/pkg/builder"
)

var tt = struct {
	xmlurl  string
	htmlurl string
	title   string
}{
	xmlurl:  "https://www.lemonde.fr/rss/une.xml",
	htmlurl: "https://www.lemonde.fr",
	title:   "Le Monde.fr - Actualités et Infos en France et dans le monde",
}

func TestNewDirectFeed(t *testing.T) {
	feed, err := builder.NewFeed(tt.xmlurl, nil)
	require.Nil(t, err)
	require.NotNil(t, feed)
	require.NotEmpty(t, feed.ID)
	require.Empty(t, feed.Tags)
	require.Equal(t, tt.xmlurl, feed.XMLURL)
	require.Equal(t, tt.title, feed.Title)
}

func TestNewIndirectFeed(t *testing.T) {
	feed, err := builder.NewFeed(tt.htmlurl, nil)
	require.Nil(t, err)
	require.NotNil(t, feed)
	require.NotEmpty(t, feed.ID)
	require.Empty(t, feed.Tags)
	require.Equal(t, tt.title, feed.Title)
}

func TestBadNewFeed(t *testing.T) {
	url := "https://keeper.nunux.org/doc/en/"
	_, err := builder.NewFeed(url, nil)
	require.NotNil(t, err)
}

func TestNewFeedWithHubAndTags(t *testing.T) {
	url := "https://medium.com/feed/netflix-techblog"
	tags := "foo,/bar_barè,/foo"
	feed, err := builder.NewFeed(url, &tags)
	require.Nil(t, err)
	require.NotNil(t, feed)
	require.NotEmpty(t, feed.ID)
	require.Equal(t, url, feed.XMLURL)
	require.Equal(t, "http://medium.superfeedr.com", *feed.HubURL)
	require.Equal(t, 2, len(feed.Tags))
	require.Equal(t, "foo", feed.Tags[0])
	require.Equal(t, "bar_barè", feed.Tags[1])
}

func TestJoinTags(t *testing.T) {
	require.Equal(t, "foo,bar", builder.JoinTags("foo", "bar"))
	require.Equal(t, "bar,foo", builder.JoinTags("", "bar", "foo", ""))
}
