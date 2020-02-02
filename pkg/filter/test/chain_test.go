package test

import (
	"strings"
	"testing"

	"github.com/ncarlier/feedpushr/v2/pkg/assert"
	"github.com/ncarlier/feedpushr/v2/pkg/builder"
	"github.com/ncarlier/feedpushr/v2/pkg/filter"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

func buildChainFilter(t *testing.T, URIs ...string) *filter.Chain {
	chain := filter.NewChainFilter()
	for _, URI := range URIs {
		condition := ""
		if strings.Contains(URI, "|") {
			parts := strings.Split(URI, "|")
			URI = parts[0]
			condition = parts[1]
		}
		filter := builder.NewFilterBuilder().FromURI(URI).Condition(&condition).Build()
		chain.Add(filter)
	}
	return chain
}

func TestNewFilterChain(t *testing.T) {
	chain := buildChainFilter(
		t,
		"title://?prefix=Hello",
		"title://?prefix=[ignore]|\"foo\" in Tags",
		"title://?prefix=[test]|\"test\" in Tags",
	)

	defs := chain.GetFilterDefs()
	assert.Equal(t, 3, len(defs), "invalid filter chain definitions")
	assert.Equal(t, "title", defs[0].Name, "invalid filter name")
	assert.Equal(t, "Hello", defs[0].Props["prefix"], "invalid filter parameter")

	article := &model.Article{
		Title: "World",
		Tags:  []string{"test"},
	}
	err := chain.Apply(article)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, "[test] Hello World", article.Title, "invalid article title")

	article = &model.Article{
		Title: "Other",
	}
	err = chain.Apply(article)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, "Hello Other", article.Title, "invalid article title")
}

func TestFilterChainCRUD(t *testing.T) {
	// CREATE
	chain := buildChainFilter(
		t,
		"title://?prefix=Hello",
	)

	defs := chain.GetFilterDefs()
	assert.Equal(t, 1, len(defs), "invalid filter chain definitions")
	_filter := defs[0]
	assert.Equal(t, "title", _filter.Name, "invalid filter type")
	assert.Equal(t, "Hello", _filter.Props["prefix"], "invalid filter property")

	// UPDATE
	update := builder.NewFilterBuilder().ID(_filter.ID).Spec(_filter.Name).Build()
	update.Props["prefix"] = "Updated"
	_, err := chain.Update(update)
	assert.Nil(t, err, "error should be nil")
	_filter = chain.GetFilterDefs()[0]
	id := _filter.ID
	assert.Equal(t, "title", _filter.Name, "invalid filter type")
	assert.Equal(t, "Updated", _filter.Props["prefix"], "invalid filter property")

	// ADD
	add := builder.NewFilterBuilder().FromURI("minify://").Build()
	_, err = chain.Add(add)
	assert.Nil(t, err, "error should be nil")
	defs = chain.GetFilterDefs()
	assert.Equal(t, 2, len(defs), "invalid filter chain definitions")
	_filter = defs[1]
	assert.Equal(t, "minify", _filter.Name, "invalid filter type")

	// DELETE
	err = chain.Remove(&model.FilterDef{ID: id})
	assert.Nil(t, err, "error should be nil")
	defs = chain.GetFilterDefs()
	assert.Equal(t, 1, len(defs), "invalid filter chain specifications")
	_filter = defs[0]
	assert.Equal(t, "minify", _filter.Name, "invalid filter type")
}
