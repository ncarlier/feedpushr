package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/filter"
	"github.com/ncarlier/feedpushr/pkg/model"
)

func buildChainFilter(t *testing.T, URIs ...string) *filter.Chain {
	chain := filter.NewChainFilter()
	for _, URI := range URIs {
		filter, err := builder.NewFilterFromURI(URI)
		assert.Nil(t, err, "error should be nil")
		chain.Add(filter)
	}
	return chain
}

func TestNewFilterChain(t *testing.T) {
	chain := buildChainFilter(
		t,
		"title://?prefix=Hello#foo,/bar,bar",
		"title://?prefix=Ignore#foo,/bar,missing",
		"title://?prefix=[test]",
	)

	defs := chain.GetFilterDefs()
	assert.Equal(t, 3, len(defs), "invalid filter chain definitions")
	assert.Equal(t, "title", defs[0].Name, "invalid filter name")
	assert.Equal(t, "Hello", defs[0].Props["prefix"], "invalid filter parameter")
	assert.Equal(t, 2, len(defs[0].Tags), "invalid filter tags")
	assert.Equal(t, "foo", defs[0].Tags[0], "invalid filter tag")

	article := &model.Article{
		Title: "World",
		Tags:  []string{"bar", "foo"},
	}
	err := chain.Apply(article)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, "[test] Hello World", article.Title, "invalid article title")

	article = &model.Article{
		Title: "Other",
	}
	err = chain.Apply(article)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, "[test] Other", article.Title, "invalid article title")
}

func TestFilterChainCRUD(t *testing.T) {
	// CREATE
	chain := buildChainFilter(
		t,
		"title://?prefix=Hello#foo,/bar,bar",
	)

	defs := chain.GetFilterDefs()
	assert.Equal(t, 1, len(defs), "invalid filter chain definitions")
	_filter := defs[0]
	assert.Equal(t, "title", _filter.Name, "invalid filter type")
	assert.Equal(t, "Hello", _filter.Props["prefix"], "invalid filter property")

	// UPDATE
	update := &app.Filter{
		ID:    _filter.ID,
		Name:  _filter.Name,
		Props: make(model.FilterProps),
		Tags:  []string{"test"},
	}
	update.Props["prefix"] = "Updated"
	_, err := chain.Update(update)
	assert.Nil(t, err, "error should be nil")
	_filter = chain.GetFilterDefs()[0]
	id := _filter.ID
	assert.Equal(t, "title", _filter.Name, "invalid filter type")
	assert.Equal(t, "Updated", _filter.Props["prefix"], "invalid filter property")
	assert.Equal(t, 1, len(_filter.Tags), "invalid filter tags")
	assert.Equal(t, "test", _filter.Tags[0], "invalid filter tag")

	// ADD
	add, err := builder.NewFilterFromURI("minify://")
	assert.Nil(t, err, "error should be nil")
	_, err = chain.Add(add)
	assert.Nil(t, err, "error should be nil")
	defs = chain.GetFilterDefs()
	assert.Equal(t, 2, len(defs), "invalid filter chain definitions")
	_filter = defs[1]
	assert.Equal(t, "minify", _filter.Name, "invalid filter type")

	// DELETE
	err = chain.Remove(&app.Filter{ID: id})
	assert.Nil(t, err, "error should be nil")
	defs = chain.GetFilterDefs()
	assert.Equal(t, 1, len(defs), "invalid filter chain specifications")
	_filter = defs[0]
	assert.Equal(t, "minify", _filter.Name, "invalid filter type")
}
