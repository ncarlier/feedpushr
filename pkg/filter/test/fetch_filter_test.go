package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/model"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/filter"
)

func TestFetchFilter(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	chain, err := filter.LoadChainFilter(db)
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, chain, "chain should not be nil")

	err = chain.AddURI("fetch://")
	assert.Nil(t, err, "error should be nil")

	link := "https://github.com/ncarlier/feedpushr"
	article := &model.Article{
		Link: link,
		Meta: make(map[string]interface{}),
	}
	err = chain.Apply(article)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, "ncarlier/feedpushr", article.Title, "invalid article title")
	assert.Equal(t, link, article.Link, "invalid article link")
	assert.Equal(t, "A simple feed aggregator daemon with sugar on top. - ncarlier/feedpushr", article.Meta["text"], "invalid description")
}
