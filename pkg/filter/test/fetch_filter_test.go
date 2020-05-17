package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/v3/pkg/assert"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

func TestFetchFilter(t *testing.T) {
	chain := buildChainFilter(t, "fetch://")

	link := "https://github.com/ncarlier/feedpushr"
	article := &model.Article{
		Link: link,
		Meta: make(map[string]interface{}),
	}
	err := chain.Apply(article)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, "ncarlier/feedpushr", article.Title, "invalid article title")
	assert.Equal(t, link, article.Link, "invalid article link")
	assert.Equal(t, "A simple feed aggregator daemon with sugar on top. - ncarlier/feedpushr", article.Text, "invalid description")
}
