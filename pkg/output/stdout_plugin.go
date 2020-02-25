package output

import (
	"fmt"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v2/pkg/expr"
	"github.com/ncarlier/feedpushr/v2/pkg/format"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

var stdoutSpec = model.Spec{
	Name: "stdout",
	Desc: "New articles are sent to the process standard output.\n\nYou can customize the payload using the [template engine](https://github.com/ncarlier/feedpushr#output-format).",
	PropsSpec: []model.PropSpec{
		{
			Name: "format",
			Desc: "Payload format (internal JSON format if not provided)",
			Type: model.Textarea,
		},
	},
}

// StdoutOutputPlugin is the STDOUT output plugin
type StdoutOutputPlugin struct{}

// Spec returns plugin spec
func (p *StdoutOutputPlugin) Spec() model.Spec {
	return stdoutSpec
}

// Build creates output provider instance
func (p *StdoutOutputPlugin) Build(output *model.OutputDef) (model.OutputProvider, error) {
	condition, err := expr.NewConditionalExpression(output.Condition)
	if err != nil {
		return nil, err
	}

	formatter, err := format.NewOutputFormatter(output)
	if err != nil {
		return nil, err
	}

	return &StdOutputProvider{
		id:        output.ID,
		alias:     output.Alias,
		spec:      stdoutSpec,
		condition: condition,
		enabled:   output.Enabled,
		formatter: formatter,
	}, nil
}

// StdOutputProvider STDOUT output provider
type StdOutputProvider struct {
	id        int
	alias     string
	spec      model.Spec
	condition *expr.ConditionalExpression
	nbSuccess uint64
	nbError   uint64
	enabled   bool
	formatter format.Formatter
}

// Send article to STDOUT.
func (op *StdOutputProvider) Send(article *model.Article) error {
	if !op.enabled || !op.condition.Match(article) {
		// Ignore if disabled or if the article doesn't match the condition
		return nil
	}
	b, err := op.formatter.Format(article)
	if err != nil {
		atomic.AddUint64(&op.nbError, 1)
		return err
	}
	fmt.Println(b.String())
	atomic.AddUint64(&op.nbSuccess, 1)
	return nil
}

// GetDef return output provider definition
func (op *StdOutputProvider) GetDef() model.OutputDef {
	result := model.OutputDef{
		ID:        op.id,
		Alias:     op.alias,
		Spec:      op.spec,
		Condition: op.condition.String(),
		Enabled:   op.enabled,
	}
	result.Props = map[string]interface{}{
		"nbSuccess": op.nbSuccess,
		"nbError":   op.nbError,
		"format":    op.formatter.Value(),
	}
	return result
}
