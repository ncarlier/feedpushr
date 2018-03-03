package builder_test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/builder"
)

func TestNewFeed(t *testing.T) {
	url := "https://keeper.nunux.org/index.xml"
	feed, err := builder.NewFeed(url)
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, feed, "feed shouldn't be nil")
	assert.NotEqual(t, "", feed.ID, "ID shouldn't be empty")
	assert.Equal(t, url, feed.XMLURL, "URL should be equals")
	assert.Equal(t, "Nunux Keeper", feed.Title, "title missmatch")
}

func TestBadNewFeed(t *testing.T) {
	url := "https://keeper.nunux.org/"
	_, err := builder.NewFeed(url)
	assert.NotNil(t, err, "error shouldn't be nil")
}

func TestNewFeedWithHub(t *testing.T) {
	url := "https://medium.com/feed/netflix-techblog"
	feed, err := builder.NewFeed(url)
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, feed, "feed shouldn't be nil")
	assert.NotEqual(t, "", feed.ID, "ID shouldn't be empty")
	assert.Equal(t, url, feed.XMLURL, "URL should be equals")
	assert.Equal(t, "http://medium.superfeedr.com", *feed.HubURL, "Hub URL should be equals")
}
