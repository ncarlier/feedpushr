package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/model"
)

const jsonFormatDesc = `
JSON Format:
{
	title: "Article title",
	description: "Article description",
	content: "Article HTML content",
	link: "Article URL",
	updated: "Article update date (String format)",
	updatedParsed: "Article update date (Date format)",
	published: "Article publication date (String format)",
	publishedParsed: "Article publication date (Date format)",
	guid: "Article feed GUID",
	meta: {
		"key": "Metadata keys and values"
	},
	tags: ["list", "of", "tags"]
}
`

// StdOutputProvider STDOUT output provider
type StdOutputProvider struct {
	name      string
	desc      string
	tags      []string
	nbSuccess uint64
}

func newStdOutputProvider(output *app.Output) *StdOutputProvider {
	return &StdOutputProvider{
		name: "stdout",
		desc: "New articles are sent as JSON document to the standard output of the process.\n\n" + jsonFormatDesc,
		tags: output.Tags,
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
		Tags: op.tags,
	}
	result.Props = map[string]interface{}{
		"nbSuccess": op.nbSuccess,
	}
	return result
}
