package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
	assert.Nil(t, err)
	assert.Equal(t, "GitHub - ncarlier/feedpushr: A simple feed aggregator daemon with sugar on top.", article.Title)
	assert.Equal(t, link, article.Link)
	assert.Equal(t, "A simple feed aggregator daemon with sugar on top. - ncarlier/feedpushr", article.Text)
}
