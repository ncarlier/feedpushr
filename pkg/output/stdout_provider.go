package output

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/mmcdole/gofeed"
)

// StdOutputProvider STDOUT output provider
type StdOutputProvider struct{}

func newStdOutputProvider() *StdOutputProvider {
	return &StdOutputProvider{}
}

// Send article to STDOUT.
func (op *StdOutputProvider) Send(article *gofeed.Item) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(article)
	fmt.Println(b.String())
	return nil
}
