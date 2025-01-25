package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/explore"
)

func TestSearchURL(t *testing.T) {
	explorer, err := explore.NewExplorer("default")
	assert.Nil(t, err)
	results, err := explorer.Search("https://www.lemonde.fr")
	assert.Nil(t, err)
	res := *results
	assert.NotEmpty(t, res)
	assert.Equal(t, "https://www.lemonde.fr/rss/une.xml", res[0].XMLURL)
}

func TestSearchQuery(t *testing.T) {
	// siked due to rsssearchhub issue
	t.SkipNow()
	explorer, err := explore.NewExplorer("default")
	assert.Nil(t, err)
	results, err := explorer.Search("tech blog")
	assert.Nil(t, err)
	assert.NotEmpty(t, *results)
}
