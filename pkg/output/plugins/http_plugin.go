package plugins

import (
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/expr"
	"github.com/ncarlier/feedpushr/v2/pkg/format"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

var supportedContentTypes = map[string]string{
	common.ContentTypeJSON: "JSON",
	common.ContentTypeText: "Text",
}

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
func (p *HTTPOutputPlugin) Build(output *model.OutputDef) (model.Output, error) {
	condition, err := expr.NewConditionalExpression(output.Condition)
	if err != nil {
		return nil, err
	}
	u, ok := output.Props["url"]
	if !ok {
		return nil, fmt.Errorf("missing URL property")
	}
	_url, err := url.ParseRequestURI(fmt.Sprintf("%v", u))
	if err != nil {
		return nil, fmt.Errorf("invalid URL property: %s", err.Error())
	}
	formatter, err := format.NewOutputFormatter(output)
	if err != nil {
		return nil, err
	}
	contentType := common.ContentTypeJSON
	if val, ok := output.Props["contentType"]; ok {
		_contentType := fmt.Sprintf("%v", val)
		if _, supported := supportedContentTypes[_contentType]; supported {
			contentType = _contentType
		}
	}

	return &HTTPOutputProvider{
		id:          output.ID,
		alias:       output.Alias,
		spec:        httpSpec,
		condition:   condition,
		enabled:     output.Enabled,
		targetURL:   _url.String(),
		contentType: contentType,
		formatter:   formatter,
	}, nil
}

// HTTPOutputProvider HTTP output provider
type HTTPOutputProvider struct {
	id          string
	alias       string
	spec        model.Spec
	condition   *expr.ConditionalExpression
	nbError     uint64
	nbSuccess   uint64
	enabled     bool
	targetURL   string
	contentType string
	formatter   format.Formatter
}

// Send article to HTTP endpoint.
func (op *HTTPOutputProvider) Send(article *model.Article) error {
	if !op.enabled || !op.condition.Match(article) {
		// Ignore if disabled or if the article doesn't match the condition
		return nil
	}

	b, err := op.formatter.Format(article)
	if err != nil {
		atomic.AddUint64(&op.nbError, 1)
		return err
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
		ID:        op.id,
		Alias:     op.alias,
		Spec:      op.spec,
		Condition: op.condition.String(),
		Enabled:   op.enabled,
		NbSuccess: op.nbSuccess,
		NbError:   op.nbError,
	}
	result.Props = map[string]interface{}{
		"url":         op.targetURL,
		"format":      op.formatter.Value(),
		"contentType": op.contentType,
	}
	return result
}
