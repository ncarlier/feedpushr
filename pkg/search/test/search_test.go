package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/v2/pkg/assert"
	"github.com/ncarlier/feedpushr/v2/pkg/search"
)

func TestSearchURL(t *testing.T) {
	engine, err := search.NewSearchEngine("default")
	assert.Nil(t, err, "error should be nil")
	results, err := engine.Search("https://keeper.nunux.org")
	assert.Nil(t, err, "error should be nil")
	res := *results
	assert.Equal(t, 1, len(res), "Results should not be empty")
	assert.Equal(t, "https://keeper.nunux.org/index.xml", res[0].XMLURL, "Results should not be empty")
}

func TestSearchQuery(t *testing.T) {
	engine, err := search.NewSearchEngine("default")
	assert.Nil(t, err, "error should be nil")
	results, err := engine.Search("tech blog")
	assert.Nil(t, err, "error should be nil")
	assert.True(t, len(*results) > 0, "Results should not be empty")
}
