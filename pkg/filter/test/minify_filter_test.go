package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
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
	assert.Nil(t, err)
	assert.Equal(t, expected, article.Content, "invalid article content")
}
