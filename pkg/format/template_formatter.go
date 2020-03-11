package format

import (
	"bytes"
	"text/template"

	"github.com/ncarlier/feedpushr/v2/pkg/format/fn"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

type templateFormatter struct {
	key   string
	value string
	tpl   *template.Template
}

// NewTemplateFormatter create new template formatter
func NewTemplateFormatter(key, format string) (Formatter, error) {
	tpl, err := template.New(key).Funcs(fn.Functions).Parse(format)
	if err != nil {
		return nil, err
	}
	return &templateFormatter{
		key:   key,
		value: format,
		tpl:   tpl,
	}, nil
}

func (f *templateFormatter) Format(article *model.Article) (*bytes.Buffer, error) {
	b := new(bytes.Buffer)
	err := f.tpl.Execute(b, article)
	return b, err
}

func (f *templateFormatter) Value() string {
	return f.value
}
