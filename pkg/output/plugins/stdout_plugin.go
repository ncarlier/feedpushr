package plugins

import (
	"fmt"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v3/pkg/format"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
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
func (p *StdoutOutputPlugin) Build(def *model.OutputDef) (model.Output, error) {
	formatter, err := format.NewOutputFormatter(def)
	if err != nil {
		return nil, err
	}
	definition := *def
	definition.Spec = stdoutSpec

	return &StdOutputProvider{
		definition: definition,
		formatter:  formatter,
	}, nil
}

// StdOutputProvider STDOUT output provider
type StdOutputProvider struct {
	definition model.OutputDef
	formatter  format.Formatter
}

// Send article to STDOUT.
func (op *StdOutputProvider) Send(article *model.Article) (bool, error) {
	b, err := op.formatter.Format(article)
	if err != nil {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, err
	}
	fmt.Println(b.String())
	atomic.AddUint64(&op.definition.NbSuccess, 1)
	return true, nil
}

// GetDef return output provider definition
func (op *StdOutputProvider) GetDef() model.OutputDef {
	return op.definition
}
