package builder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/builder"
)

const feedTitle = "Le Monde.fr - Actualités et Infos en France et dans le monde"

func TestNewDirectFeed(t *testing.T) {
	url := "https://www.lemonde.fr/rss/une.xml"
	feed, err := builder.NewFeed(url, nil)
	assert.Nil(t, err)
	assert.NotNil(t, feed)
	assert.NotEmpty(t, feed.ID)
	assert.Empty(t, feed.Tags)
	assert.Equal(t, url, feed.XMLURL)
	assert.Equal(t, feedTitle, feed.Title)
}

func TestNewIndirectFeed(t *testing.T) {
	url := "https://www.lemonde.fr"
	feed, err := builder.NewFeed(url, nil)
	assert.Nil(t, err)
	assert.NotNil(t, feed)
	assert.NotEmpty(t, feed.ID)
	assert.Empty(t, feed.Tags)
	assert.Equal(t, feedTitle, feed.Title)
}

func TestBadNewFeed(t *testing.T) {
	url := "https://keeper.nunux.org/doc/en/"
	_, err := builder.NewFeed(url, nil)
	assert.NotNil(t, err)
}

func TestNewFeedWithHubAndTags(t *testing.T) {
	url := "https://medium.com/feed/netflix-techblog"
	tags := "foo,/bar_barè,/foo"
	feed, err := builder.NewFeed(url, &tags)
	assert.Nil(t, err)
	assert.NotNil(t, feed)
	assert.NotEmpty(t, feed.ID)
	assert.Equal(t, url, feed.XMLURL)
	assert.Equal(t, "http://medium.superfeedr.com", *feed.HubURL)
	assert.Equal(t, 2, len(feed.Tags))
	assert.Equal(t, "foo", feed.Tags[0])
	assert.Equal(t, "bar_barè", feed.Tags[1])
}

func TestJoinTags(t *testing.T) {
	assert.Equal(t, "foo,bar", builder.JoinTags("foo", "bar"))
	assert.Equal(t, "bar,foo", builder.JoinTags("", "bar", "foo", ""))
}
