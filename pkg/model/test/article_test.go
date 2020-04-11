package test

import (
	"fmt"
	"testing"

	"github.com/ncarlier/feedpushr/v3/pkg/assert"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

var testCases = []struct {
	articleTags []string
	tags        []string
	result      bool
}{
	{
		[]string{"foo", "bar"},
		[]string{"foo"},
		true,
	},
	{
		[]string{"foo", "bar"},
		[]string{"foo", "bar"},
		true,
	},
	{
		[]string{"foo", "bar"},
		[]string{"foo", "bar", "bip"},
		false,
	},
	{
		[]string{"foo", "bar", "bip"},
		[]string{"foo", "bar"},
		true,
	},
	{
		[]string{"foo", "bar"},
		[]string{"!foo"},
		false,
	},
	{
		[]string{"foo", "bar"},
		[]string{"foo", "!bip"},
		true,
	},
}

func TestMatchArticle(t *testing.T) {
	for idx, tc := range testCases {
		article := model.Article{Tags: tc.articleTags}
		assert.Equal(t, tc.result, article.Match(tc.tags), fmt.Sprintf("Error with test case #%d", idx))
	}
}
