package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/pkg/model"
)

// StdOutputProvider STDOUT output provider
type StdOutputProvider struct {
	name      string
	desc      string
	nbSuccess uint64
}

func newStdOutputProvider() *StdOutputProvider {
	return &StdOutputProvider{
		name: "stdout",
		desc: "New articles are sent as JSON document to the standard output of the process.",
	}
}

// Send article to STDOUT.
func (op *StdOutputProvider) Send(article *model.Article) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(article)
	fmt.Println(b.String())
	atomic.AddUint64(&op.nbSuccess, 1)
	return nil
}

// GetSpec return output provider specifications
func (op *StdOutputProvider) GetSpec() model.OutputSpec {
	result := model.OutputSpec{
		Name: op.name,
		Desc: op.desc,
	}
	result.Props = map[string]interface{}{
		"nbSuccess": op.nbSuccess,
	}
	return result
}
