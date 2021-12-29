package plugins

import (
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/format"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

var supportedContentTypes = map[string]string{
	common.ContentTypeJSON: "JSON",
	common.ContentTypeText: "Text",
}

var httpSpec = model.Spec{
	Name: "http",
	Desc: "New articles are sent to a HTTP endpoint (POST).\n\nYou can customize the payload using the [template engine](https://github.com/ncarlier/feedpushr#output-format).",
	PropsSpec: []model.PropSpec{
		{
			Name: "url",
			Desc: "Target URL",
			Type: model.URL,
		},
		{
			Name:    "contentType",
			Desc:    "Content type",
			Type:    model.Select,
			Options: supportedContentTypes,
		},
		{
			Name: "format",
			Desc: "Payload format (internal JSON format by defaut)",
			Type: model.Textarea,
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
func (p *HTTPOutputPlugin) Build(def *model.OutputDef) (model.Output, error) {
	u, ok := def.Props["url"]
	if !ok {
		return nil, fmt.Errorf("missing URL property")
	}
	targetURL, err := url.ParseRequestURI(fmt.Sprintf("%v", u))
	if err != nil {
		return nil, fmt.Errorf("invalid URL property: %s", err.Error())
	}
	formatter, err := format.NewOutputFormatter(def)
	if err != nil {
		return nil, err
	}
	contentType := common.ContentTypeJSON
	if val, ok := def.Props["contentType"]; ok {
		_contentType := fmt.Sprintf("%v", val)
		if _, supported := supportedContentTypes[_contentType]; supported {
			contentType = _contentType
		}
	}

	definition := *def
	definition.Spec = httpSpec
	definition.Props["url"] = targetURL.String()

	return &HTTPOutputProvider{
		definition:  definition,
		targetURL:   targetURL.String(),
		contentType: contentType,
		formatter:   formatter,
	}, nil
}

// HTTPOutputProvider HTTP output provider
type HTTPOutputProvider struct {
	definition  model.OutputDef
	targetURL   string
	contentType string
	formatter   format.Formatter
}

// Send article to HTTP endpoint.
func (op *HTTPOutputProvider) Send(article *model.Article) (bool, error) {
	b, err := op.formatter.Format(article)
	if err != nil {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, err
	}

	req, err := http.NewRequest("POST", op.targetURL, b)
	if err != nil {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, err
	}
	req.Header.Set("User-Agent", common.UserAgent)
	req.Header.Set("Content-Type", op.contentType)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	atomic.AddUint64(&op.definition.NbSuccess, 1)
	return true, nil
}

// GetDef return output provider definition
func (op *HTTPOutputProvider) GetDef() model.OutputDef {
	return op.definition
}
