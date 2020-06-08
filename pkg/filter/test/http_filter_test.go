package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

func TestHTTPFilter(t *testing.T) {
	chain := buildChainFilter(t, "http://?url=https://run.mocky.io/v3/64073093-4e20-412d-8e0a-a6a1e23c01bd")

	article := &model.Article{
		Title:   "hello world",
		Content: "<p>this should be replaced</p>",
		Text:    "this should be kept",
		Meta:    make(map[string]interface{}),
	}
	article.Meta["foo"] = "bar"
	err := chain.Apply(article)
	assert.Nil(t, err)
	assert.Equal(t, "A mock", article.Title)
	assert.Equal(t, "<p>with a fake content</p>", article.Content)
	assert.Equal(t, "this should be kept", article.Text)
	assert.Equal(t, "bar", article.Meta["foo"])
}
