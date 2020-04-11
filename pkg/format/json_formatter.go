package format

import (
	"bytes"
	"encoding/json"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

type jsonFormatter struct{}

// NewJSONFormatter create new JSON formatter
func NewJSONFormatter() Formatter {
	return &jsonFormatter{}
}

func (f *jsonFormatter) Format(article *model.Article) (*bytes.Buffer, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(article)
	return b, err
}

func (f *jsonFormatter) Value() string {
	return ""
}
