package format

import "bytes"

import "github.com/ncarlier/feedpushr/v3/pkg/model"

// Formatter is an interface for article formating
type Formatter interface {
	Format(article *model.Article) (*bytes.Buffer, error)
	Value() string
}
