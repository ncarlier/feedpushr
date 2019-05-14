package opml_test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/opml"
)

func TestNewOPML(t *testing.T) {
	o, err := opml.NewOPMLFromFile("./tc_simple.xml")
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, o, "OPML object shouldn't be nil")
	assert.Equal(t, "Simple test case", o.Head.Title, "Header don't match")
	assert.Equal(t, 1, len(o.Body.Outlines), "Body don't match")
	assert.Equal(t, "Hashicorp Blog", o.Body.Outlines[0].Title, "Outline title don't match")
}

func TestNewOPMLWithOutlineCategories(t *testing.T) {
	o, err := opml.NewOPMLFromFile("./tc_with_outline_categories.xml")
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, o, "OPML object shouldn't be nil")
	assert.Equal(t, "Test case with categories", o.Head.Title, "Header don't match")
	assert.Equal(t, 2, len(o.Body.Outlines), "Body don't match")
	assert.Equal(t, "Computer Science", o.Body.Outlines[0].Title, "Outline title don't match")
}

func TestNewOPMLWithInlineCategories(t *testing.T) {
	o, err := opml.NewOPMLFromFile("./tc_with_inline_categories.xml")
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, o, "OPML object shouldn't be nil")
	assert.Equal(t, "Test case with categories", o.Head.Title, "Header don't match")
	assert.Equal(t, 4, len(o.Body.Outlines), "Body don't match")
	assert.Equal(t, "/Computer Science", o.Body.Outlines[0].Category, "Outline title don't match")
}
