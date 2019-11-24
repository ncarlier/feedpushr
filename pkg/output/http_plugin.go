package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"
	"text/template"

	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/model"
)

var supportedContentTypes = []string{common.ContentTypeJSON, common.ContentTypeText}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
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
			Desc: "Payload format (internal JSON format if not provided)",
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
func (p *HTTPOutputPlugin) Build(output *model.OutputDef) (model.OutputProvider, error) {
	u, ok := output.Props["url"]
	if !ok {
		return nil, fmt.Errorf("missing URL property")
	}
	_url, err := url.ParseRequestURI(fmt.Sprintf("%v", u))
	if err != nil {
		return nil, fmt.Errorf("invalid URL property: %s", err.Error())
	}
	var tpl *template.Template
	var format string
	if _format, ok := output.Props["format"]; ok && _format != "" {
		tplName := fmt.Sprintf("stdout-%d", output.ID)
		format = fmt.Sprintf("%v", _format)
		var err error
		tpl, err = template.New(tplName).Parse(format)
		if err != nil {
			return nil, err
		}
	}
	contentType := common.ContentTypeJSON
	if ct, ok := output.Props["contentType"]; ok && contains(supportedContentTypes, fmt.Sprintf("%v", ct)) {
		contentType = fmt.Sprintf("%v", ct)
	}

	return &HTTPOutputProvider{
		id:          output.ID,
		alias:       output.Alias,
		spec:        httpSpec,
		tags:        output.Tags,
		enabled:     output.Enabled,
		targetURL:   _url.String(),
		contentType: contentType,
		format:      format,
		tpl:         tpl,
	}, nil
}

var httpOutputPlugin = &HTTPOutputPlugin{}

// HTTPOutputProvider HTTP output provider
type HTTPOutputProvider struct {
	id          int
	alias       string
	spec        model.Spec
	tags        []string
	nbError     uint64
	nbSuccess   uint64
	enabled     bool
	targetURL   string
	contentType string
	format      string
	tpl         *template.Template
}

// Send article to HTTP endpoint.
func (op *HTTPOutputProvider) Send(article *model.Article) error {
	b := new(bytes.Buffer)
	if op.tpl != nil {
		if err := op.tpl.Execute(b, article); err != nil {
			return err
		}
	} else {
		if err := json.NewEncoder(b).Encode(article); err != nil {
			return err
		}
	}

	req, err := http.NewRequest("POST", op.targetURL, b)
	if err != nil {
		atomic.AddUint64(&op.nbError, 1)
		return err
	}
	req.Header.Set("User-Agent", common.UserAgent)
	req.Header.Set("Content-Type", op.contentType)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		atomic.AddUint64(&op.nbError, 1)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
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
		"nbError":     op.nbError,
		"nbSuccess":   op.nbSuccess,
		"url":         op.targetURL,
		"format":      op.format,
		"contentType": op.contentType,
	}
	return result
}
