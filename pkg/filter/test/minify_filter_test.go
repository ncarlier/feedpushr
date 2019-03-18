package filter_test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/feedpushr/pkg/plugin"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/filter"
)

func setupMinifyTestCase(t *testing.T) *filter.Chain {
	pr := &plugin.Registry{}
	filters := []string{"minify://"}
	chain, err := filter.NewChainFilter(filters, pr)
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, chain, "chain should not be nil")
	return chain
}

func TestMinifyFilter(t *testing.T) {
	filterChain := setupMinifyTestCase(t)
	article := &model.Article{
		Content: `<ul>
			<li>
				<p>Hello World</p>
				<img />
			</li>
		</ul>`,
	}
	expected := "<ul><li><p>Hello World</p><img></ul>"
	err := filterChain.Apply(article)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, expected, article.Content, "invalid article content")
}
