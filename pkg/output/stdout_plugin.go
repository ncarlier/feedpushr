package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"text/template"

	"github.com/ncarlier/feedpushr/pkg/expr"
	"github.com/ncarlier/feedpushr/pkg/model"
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
	var tpl *template.Template
	var format string
	if _format, ok := output.Props["format"]; ok && _format != "" {
		tplName := fmt.Sprintf("stdout-%d", output.ID)
		format = fmt.Sprintf("%v", _format)
		var err error
		tpl, err = template.New(tplName).Parse(format)
		if err != nil {
			return nil, err
		}
	}
	return &StdOutputProvider{
		id:        output.ID,
		alias:     output.Alias,
		spec:      stdoutSpec,
		condition: condition,
		enabled:   output.Enabled,
		format:    format,
		tpl:       tpl,
	}, nil
}

// StdOutputProvider STDOUT output provider
type StdOutputProvider struct {
	id        int
	alias     string
	spec      model.Spec
	condition *expr.ConditionalExpression
	nbSuccess uint64
	enabled   bool
	format    string
	tpl       *template.Template
}

// Send article to STDOUT.
func (op *StdOutputProvider) Send(article *model.Article) error {
	if !op.enabled || !op.condition.Match(article) {
		// Ignore if disabled or if the article doesn't match the condition
		return nil
	}
	b := new(bytes.Buffer)
	if op.tpl != nil {
		if err := op.tpl.Execute(b, article); err != nil {
			return err
		}
	} else {
		if err := json.NewEncoder(b).Encode(article); err != nil {
			return err
		}
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
		"format":    op.format,
	}
	return result
}
