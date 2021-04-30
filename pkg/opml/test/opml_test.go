package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/opml"
)

func TestNewOPML(t *testing.T) {
	o, err := opml.NewOPMLFromFile("./tc_simple.xml")
	assert.Nil(t, err)
	assert.NotNil(t, o, "OPML object shouldn't be nil")
	assert.Equal(t, "Simple test case", o.Head.Title)
	assert.Len(t, o.Body.Outlines, 1)
	assert.Equal(t, "Hashicorp Blog", o.Body.Outlines[0].Title, "Outline title don't match")
}

func TestNewOPMLWithOutlineCategories(t *testing.T) {
	o, err := opml.NewOPMLFromFile("./tc_with_outline_categories.xml")
	assert.Nil(t, err)
	assert.NotNil(t, o, "OPML object shouldn't be nil")
	assert.Equal(t, "Test case with categories", o.Head.Title)
	assert.Len(t, o.Body.Outlines, 2)
	assert.Equal(t, "Computer Science", o.Body.Outlines[0].Title, "Outline title don't match")
}

func TestNewOPMLWithInlineCategories(t *testing.T) {
	o, err := opml.NewOPMLFromFile("./tc_with_inline_categories.xml")
	assert.Nil(t, err)
	assert.NotNil(t, o, "OPML object shouldn't be nil")
	assert.Equal(t, "Test case with categories", o.Head.Title)
	assert.Len(t, o.Body.Outlines, 3)
	assert.Equal(t, "/Computer Science", o.Body.Outlines[0].Category, "Outline title don't match")
}
