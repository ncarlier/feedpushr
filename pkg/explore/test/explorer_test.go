package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/v2/pkg/assert"
	"github.com/ncarlier/feedpushr/v2/pkg/explore"
)

func TestSearchURL(t *testing.T) {
	explorer, err := explore.NewExplorer("default")
	assert.Nil(t, err, "error should be nil")
	results, err := explorer.Search("https://keeper.nunux.org")
	assert.Nil(t, err, "error should be nil")
	res := *results
	assert.Equal(t, 1, len(res), "Results should not be empty")
	assert.Equal(t, "https://keeper.nunux.org/index.xml", res[0].XMLURL, "Results should not be empty")
}

func TestSearchQuery(t *testing.T) {
	explorer, err := explore.NewExplorer("default")
	assert.Nil(t, err, "error should be nil")
	results, err := explorer.Search("tech blog")
	assert.Nil(t, err, "error should be nil")
	assert.True(t, len(*results) > 0, "Results should not be empty")
}
