package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/model"

	"github.com/ncarlier/feedpushr/pkg/assert"
)

func TestMinifyFilter(t *testing.T) {
	chain := buildChainFilter(t, "minify://")

	article := &model.Article{
		Content: `<ul>
			<li>
				<p>Hello World</p>
				<img />
			</li>
		</ul>`,
	}
	expected := "<ul><li><p>Hello World</p><img></ul>"
	err := chain.Apply(article)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, expected, article.Content, "invalid article content")
}
