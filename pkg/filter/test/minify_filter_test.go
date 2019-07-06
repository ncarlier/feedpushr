package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/model"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/filter"
)

func TestMinifyFilter(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	chain, err := filter.LoadChainFilter(db)
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, chain, "chain should not be nil")

	err = chain.AddURI("minify://")
	assert.Nil(t, err, "error should be nil")

	article := &model.Article{
		Content: `<ul>
			<li>
				<p>Hello World</p>
				<img />
			</li>
		</ul>`,
	}
	expected := "<ul><li><p>Hello World</p><img></ul>"
	err = chain.Apply(article)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, expected, article.Content, "invalid article content")
}
