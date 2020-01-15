package expr

import (
	"fmt"
	"strings"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/rs/zerolog/log"

	"github.com/ncarlier/feedpushr/pkg/model"
)

// ConditionalExpression is a model for a conditional expression applied on an article
type ConditionalExpression struct {
	expression string
	prog       *vm.Program
}

// NewConditionalExpression creates a new conditional expression
func NewConditionalExpression(expression string) (*ConditionalExpression, error) {
	var prog *vm.Program
	var err error
	if strings.TrimSpace(expression) != "" {
		prog, err = expr.Compile(expression, expr.Env(model.Article{}))
		if err != nil {
			return nil, fmt.Errorf("invalid conditional expression: %s", err.Error())
		}
	}
	return &ConditionalExpression{
		expression: expression,
		prog:       prog,
	}, nil
}

// Match test ifthe article match the conditional expression
func (c *ConditionalExpression) Match(article *model.Article) bool {
	if c.prog == nil {
		return true
	}
	output, err := expr.Run(c.prog, article)
	if err != nil {
		log.Error().Err(err).Str("expr", c.expression).Str("article", article.Link).Msg("unable to run expression on the article")
		return false
	}
	return asBooleanValue(output)
}

// String returns the string expression
func (c *ConditionalExpression) String() string {
	return c.expression
}

func asBooleanValue(i1 interface{}) bool {
	if i1 == nil {
		return false
	}
	switch i2 := i1.(type) {
	default:
		return false
	case bool:
		return i2
	case string:
		return i2 == "true"
	case int:
		return i2 != 0
	case *bool:
		if i2 == nil {
			return false
		}
		return *i2
	case *string:
		if i2 == nil {
			return false
		}
		return *i2 == "true"
	case *int:
		if i2 == nil {
			return false
		}
		return *i2 != 0
	}
}
