package output

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ncarlier/feedpushr/pkg/model"
)

// StdOutputProvider STDOUT output provider
type StdOutputProvider struct{}

func newStdOutputProvider() *StdOutputProvider {
	return &StdOutputProvider{}
}

// Send article to STDOUT.
func (op *StdOutputProvider) Send(article *model.Article) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(article)
	fmt.Println(b.String())
	return nil
}
