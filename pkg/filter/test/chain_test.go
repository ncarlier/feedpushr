package filter_test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/model"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/filter"
)

func TestNewFilterChain(t *testing.T) {
	filters := []string{
		"title://?prefix=Hello#foo,/bar,bar",
		"title://?prefix=Ignore#foo,/bar,missing",
		"title://?prefix=[test]",
	}
	chain, err := filter.NewChainFilter(filters)
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, chain, "chain should not be nil")
	specs := chain.GetSpec()
	assert.Equal(t, 3, len(specs), "invalid filter chain specifications")
	assert.Equal(t, "title", specs[0].Name, "invalid filter name")
	assert.Equal(t, "Hello", specs[0].Props["prefix"], "invalid filter parameter")
	assert.Equal(t, 2, len(specs[0].Tags), "invalid filter tags")
	assert.Equal(t, "foo", specs[0].Tags[0], "invalid filter tag")

	article := &model.Article{
		Title: "World",
		Tags:  []string{"bar", "foo"},
	}
	err = chain.Apply(article)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, "[test] Hello World", article.Title, "invalid article title")

	article = &model.Article{
		Title: "Other",
	}
	err = chain.Apply(article)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, "[test] Other", article.Title, "invalid article title")
}
