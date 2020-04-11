package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/v3/pkg/assert"
	"github.com/ncarlier/feedpushr/v3/pkg/expr"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

func TesInvalidExpressionSyntax(t *testing.T) {
	_, err := expr.NewConditionalExpression("###")
	assert.NotNil(t, err, "expression should not be valid")
}

func TestMatchingExpression(t *testing.T) {
	condition, err := expr.NewConditionalExpression("\"foo\" in Tags")
	assert.Nil(t, err, "expression should be valid")
	assert.NotNil(t, condition, "expression should not be nil")

	article := &model.Article{
		Title: "World",
		Tags:  []string{"bar", "foo"},
	}

	ok := condition.Match(article)
	assert.True(t, ok, "article should match")
}

func TestNotMatchingExpression(t *testing.T) {
	condition, err := expr.NewConditionalExpression("len(Title) > 10")
	assert.Nil(t, err, "expression should be valid")
	assert.NotNil(t, condition, "expression should not be nil")

	article := &model.Article{
		Title: "World",
		Tags:  []string{"bar", "foo"},
	}

	ok := condition.Match(article)
	assert.True(t, !ok, "article should not match")
}
