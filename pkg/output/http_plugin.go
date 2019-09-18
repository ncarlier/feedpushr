package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/pkg/model"
)

var httpSpec = model.Spec{
	Name: "http",
	Desc: "New articles are sent as JSON document to an HTTP endpoint (POST).\n\n" + jsonFormatDesc,
	PropsSpec: []model.PropSpec{
		{
			Name: "url",
			Desc: "Target URL",
			Type: model.URL,
		},
	},
}

// HTTPOutputPlugin is the HTTP output plugin
type HTTPOutputPlugin struct{}

// Spec returns plugin spec
func (p *HTTPOutputPlugin) Spec() model.Spec {
	return httpSpec
}

// Build creates output provider instance
func (p *HTTPOutputPlugin) Build(output *model.OutputDef) (model.OutputProvider, error) {
	u, ok := output.Props["url"]
	if !ok {
		return nil, fmt.Errorf("missing URL property")
	}
	_url, err := url.ParseRequestURI(fmt.Sprintf("%v", u))
	if err != nil {
		return nil, fmt.Errorf("invalid URL property: %s", err.Error())
	}
	return &HTTPOutputProvider{
		id:        output.ID,
		alias:     output.Alias,
		spec:      httpSpec,
		tags:      output.Tags,
		targetURL: _url.String(),
		enabled:   output.Enabled,
	}, nil
}

var httpOutputPlugin = &HTTPOutputPlugin{}

// HTTPOutputProvider HTTP output provider
type HTTPOutputProvider struct {
	id        int
	alias     string
	spec      model.Spec
	tags      []string
	nbError   uint64
	nbSuccess uint64
	targetURL string
	enabled   bool
}

// Send article to HTTP endpoint.
func (op *HTTPOutputProvider) Send(article *model.Article) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(article)
	resp, err := http.Post(op.targetURL, "application/json; charset=utf-8", b)
	if err != nil {
		atomic.AddUint64(&op.nbError, 1)
		return err
	} else if resp.StatusCode >= 300 {
		atomic.AddUint64(&op.nbError, 1)
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	atomic.AddUint64(&op.nbSuccess, 1)
	return nil
}

// GetDef return output provider definition
func (op *HTTPOutputProvider) GetDef() model.OutputDef {
	result := model.OutputDef{
		ID:      op.id,
		Alias:   op.alias,
		Spec:    op.spec,
		Tags:    op.tags,
		Enabled: op.enabled,
	}
	result.Props = map[string]interface{}{
		"url":       op.targetURL,
		"nbError":   op.nbError,
		"nbSuccess": op.nbSuccess,
	}
	return result
}
