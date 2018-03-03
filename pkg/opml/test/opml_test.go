package opml_test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/opml"
)

func TestNewOPML(t *testing.T) {
	o, err := opml.NewOPMLFromFile("./sample.xml")
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, o, "OPML object shouldn't be nil")
	assert.Equal(t, "Subscriptions sample.", o.Head.Title, "Header don't match")
	assert.Equal(t, 1, len(o.Body.Outlines), "Body don't match")
	assert.Equal(t, "Hashicorp Blog", o.Body.Outlines[0].Title, "Outline title don't match")
}
