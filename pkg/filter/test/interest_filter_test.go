package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

func TestInterestFilter(t *testing.T) {
	chain := buildChainFilter(t, "interest://?wordlist=news,dolor sito")

	article := &model.Article{
		Title: "this is an example",
		Text:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	}
	err := chain.Apply(article)
	assert.NotNil(t, err)
	assert.Equal(t, err, common.ErrArticleShouldBeIgnored)

	article.Title = "Breaking news"
	err = chain.Apply(article)
	assert.Nil(t, err)
}
