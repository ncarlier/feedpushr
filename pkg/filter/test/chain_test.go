package test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/filter"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

func buildChainFilter(t *testing.T, URIs ...string) *filter.Chain {
	chain, _ := filter.NewChainFilter(model.FilterDefCollection{})
	for _, URI := range URIs {
		condition := ""
		if strings.Contains(URI, "|") {
			parts := strings.Split(URI, "|")
			URI = parts[0]
			condition = parts[1]
		}
		filter := filter.NewBuilder().FromURI(URI).Condition(&condition).Build()
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
	assert.Nil(t, err)
	assert.Equal(t, "[test] Hello World", article.Title)

	article = &model.Article{
		Title: "Other",
	}
	err = chain.Apply(article)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Other", article.Title)
}

func TestFilterChainCRUD(t *testing.T) {
	// CREATE
	chain := buildChainFilter(
		t,
		"title://?prefix=Hello",
	)

	defs := chain.GetFilterDefs()
	assert.Len(t, defs, 1, "invalid filter chain definitions")
	_filter := defs[0]
	assert.Equal(t, "title", _filter.Name, "invalid filter name")
	assert.Equal(t, "Hello", _filter.Props["prefix"], "invalid filter property")
	assert.NotEmpty(t, _filter.ID)
	id := _filter.ID

	// UPDATE
	update := filter.NewBuilder().Spec(_filter.Name).Build()
	update.Props["prefix"] = "Updated"
	_, err := chain.Update(id, update)
	assert.Nil(t, err)
	_filter = chain.GetFilterDefs()[0]
	assert.Equal(t, "title", _filter.Name, "invalid filter type")
	assert.Equal(t, "Updated", _filter.Props["prefix"], "invalid filter property")

	// ADD
	add := filter.NewBuilder().FromURI("minify://").Build()
	_, err = chain.Add(add)
	assert.Nil(t, err)
	defs = chain.GetFilterDefs()
	assert.Len(t, defs, 2, "invalid filter chain definitions")
	_filter = defs[1]
	assert.Equal(t, "minify", _filter.Name, "invalid filter type")

	// DELETE
	err = chain.Remove(id)
	assert.Nil(t, err)
	defs = chain.GetFilterDefs()
	assert.Len(t, defs, 1, "invalid filter chain specifications")
	_filter = defs[0]
	assert.Equal(t, "minify", _filter.Name, "invalid filter type")
}
