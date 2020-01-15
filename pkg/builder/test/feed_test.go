package builder_test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/builder"
)

func TestNewDirectFeed(t *testing.T) {
	url := "https://keeper.nunux.org/index.xml"
	feed, err := builder.NewFeed(url, nil)
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, feed, "feed shouldn't be nil")
	assert.NotEqual(t, "", feed.ID, "ID shouldn't be empty")
	assert.Equal(t, 0, len(feed.Tags), "Tags should be empty")
	assert.Equal(t, url, feed.XMLURL, "URL should be equals")
	assert.Equal(t, "Nunux Keeper", feed.Title, "title missmatch")
}

func TestNewIndirectFeed(t *testing.T) {
	url := "https://keeper.nunux.org"
	feed, err := builder.NewFeed(url, nil)
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, feed, "feed shouldn't be nil")
	assert.NotEqual(t, "", feed.ID, "ID shouldn't be empty")
	assert.Equal(t, 0, len(feed.Tags), "Tags should be empty")
	assert.Equal(t, "Nunux Keeper", feed.Title, "title missmatch")
}

func TestBadNewFeed(t *testing.T) {
	url := "https://keeper.nunux.org/doc/en/"
	_, err := builder.NewFeed(url, nil)
	assert.NotNil(t, err, "error shouldn't be nil")
}

func TestNewFeedWithHubAndTags(t *testing.T) {
	url := "https://medium.com/feed/netflix-techblog"
	tags := "foo,/bar_barè,/foo"
	feed, err := builder.NewFeed(url, &tags)
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, feed, "feed shouldn't be nil")
	assert.NotEqual(t, "", feed.ID, "ID shouldn't be empty")
	assert.Equal(t, url, feed.XMLURL, "URL should be equals")
	assert.Equal(t, "http://medium.superfeedr.com", *feed.HubURL, "Hub URL should be equals")
	assert.Equal(t, 2, len(feed.Tags), "Tags should not be empty")
	assert.Equal(t, "foo", feed.Tags[0], "Tag should be equals")
	assert.Equal(t, "bar_barè", feed.Tags[1], "Tag should be equals")
}

func TestJoinTags(t *testing.T) {
	assert.Equal(t, "foo,bar", builder.JoinTags("foo", "bar"), "")
	assert.Equal(t, "bar,foo", builder.JoinTags("", "bar", "foo", ""), "")
}
