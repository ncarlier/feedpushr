package expr

import (
	"fmt"
	"strings"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

var exprPlugins = map[string]interface{}{
	"toLower": strings.ToLower,
	"toUpper": strings.ToUpper,
}

// ConditionalExpression is a model for a conditional expression applied on an article
type ConditionalExpression struct {
	expression string
	prog       *vm.Program
}

// NewConditionalExpression creates a new conditional expression
func NewConditionalExpression(expression string) (*ConditionalExpression, error) {
	var prog *vm.Program
	if strings.TrimSpace(expression) != "" {
		args, err := buildExprArgs(model.Article{})
		if err != nil {
			return nil, fmt.Errorf("unable to build expression arguments: %s", err.Error())
		}

		options := []expr.Option{
			expr.Env(args),
			expr.AsBool(),
		}
		prog, err = expr.Compile(expression, options...)
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
	args, err := buildExprArgs(article)
	if err != nil {
		log.Error().Err(err).Str("expr", c.expression).Str("article", article.Link).Msg("unable to build expression arguments")
		return false
	}
	output, err := expr.Run(c.prog, args)
	if err != nil {
		log.Error().Err(err).Str("expr", c.expression).Str("article", article.Link).Msg("unable to run expression on the article")
		return false
	}
	return output.(bool)
}

// String returns the string expression
func (c *ConditionalExpression) String() string {
	return c.expression
}

func buildExprArgs(obj interface{}) (map[string]interface{}, error) {
	env := map[string]interface{}{}
	if err := mapstructure.Decode(obj, &env); err != nil {
		return nil, err
	}
	for k, v := range exprPlugins {
		env[k] = v
	}
	return env, nil
}
