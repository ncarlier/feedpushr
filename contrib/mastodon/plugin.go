package main

import (
	"fmt"
	"net/url"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v2/pkg/expr"
	"github.com/ncarlier/feedpushr/v2/pkg/format"
	"github.com/ncarlier/feedpushr/v2/pkg/format/fn"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

var tootVisibilities = map[string]string{
	"public":   "Public",
	"private":  "Private",
	"direct":   "Direct",
	"unlisted": "Unlisted",
}

var spec = model.Spec{
	Name: "mastodon",
	Desc: "Send new articles as *Toot* to a Mastodon instance.",
	PropsSpec: []model.PropSpec{
		{
			Name: "url",
			Desc: "Target URL",
			Type: model.Text,
		},
		{
			Name: "token",
			Desc: "Access token",
			Type: model.Password,
		},
		{
			Name:    "visibility",
			Desc:    "Toot visibility",
			Type:    model.Select,
			Options: tootVisibilities,
		},
		{
			Name: "format",
			Desc: "Toot format (default: `{{.Title}}\\n{{.Link}}`)",
			Type: model.Textarea,
		},
	},
}

// MastodonOutputPlugin is the Mastodon output plugin
type MastodonOutputPlugin struct{}

// Spec returns plugin spec
func (p *MastodonOutputPlugin) Spec() model.Spec {
	return spec
}

// Build creates Mastodon output provider instance
func (p *MastodonOutputPlugin) Build(output *model.OutputDef) (model.Output, error) {
	condition, err := expr.NewConditionalExpression(output.Condition)
	if err != nil {
		return nil, err
	}
	// Default format
	if frmt, ok := output.Props["format"]; !ok || frmt == "" {
		output.Props["format"] = "{{.Title}}\n{{.Link}}"
	}
	formatter, err := format.NewOutputFormatter(output)
	if err != nil {
		return nil, err
	}
	u := output.Props.Get("url")
	if u == "" {
		return nil, fmt.Errorf("missing URL property")
	}
	_url, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, fmt.Errorf("invalid URL property: %s", err.Error())
	}
	_url.Path = "/api/v1/statuses"
	accessToken := output.Props.Get("token")
	if accessToken == "" {
		return nil, fmt.Errorf("missing access token property")
	}
	visibility := output.Props.Get("visibility")
	if _, exists := tootVisibilities[visibility]; !exists {
		visibility = "public"
	}
	return &MastodonOutputProvider{
		id:          output.ID,
		alias:       output.Alias,
		spec:        spec,
		condition:   condition,
		formatter:   formatter,
		enabled:     output.Enabled,
		targetURL:   _url.String(),
		accessToken: accessToken,
		visibility:  visibility,
	}, nil
}

// MastodonOutputProvider output provider to send articles to Mastodon
type MastodonOutputProvider struct {
	id          string
	alias       string
	spec        model.Spec
	condition   *expr.ConditionalExpression
	formatter   format.Formatter
	enabled     bool
	nbError     uint64
	nbSuccess   uint64
	targetURL   string
	accessToken string
	visibility  string
}

// Send article to a Mastodon instance.
func (op *MastodonOutputProvider) Send(article *model.Article) error {
	if !op.enabled || !op.condition.Match(article) {
		// Ignore if disabled or if the article doesn't match the condition
		return nil
	}
	b, err := op.formatter.Format(article)
	if err != nil {
		atomic.AddUint64(&op.nbError, 1)
		return err
	}
	toot := Toot{
		Status:     fn.Truncate(500, b.String()),
		Sensitive:  false,
		Visibility: op.visibility,
	}
	if err := sendToMastodon(toot, op.targetURL, op.accessToken); err != nil {
		atomic.AddUint64(&op.nbError, 1)
		return err
	}
	atomic.AddUint64(&op.nbSuccess, 1)
	return nil
}

// GetDef return output definition
func (op *MastodonOutputProvider) GetDef() model.OutputDef {
	result := model.OutputDef{
		ID:        op.id,
		Alias:     op.alias,
		Spec:      op.spec,
		Condition: op.condition.String(),
		Enabled:   op.enabled,
	}
	result.Props = map[string]interface{}{
		"url":        op.targetURL,
		"token":      op.accessToken,
		"visibility": op.visibility,
		"nbError":    op.nbError,
		"nbSuccess":  op.nbSuccess,
		"format":     op.formatter.Value(),
	}
	return result
}

// GetPluginSpec return plugin informations
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: spec,
		Type: model.OutputPluginType,
	}
}

// GetOutputPlugin returns output plugin
func GetOutputPlugin() (op model.OutputPlugin, err error) {
	return &MastodonOutputPlugin{}, nil
}
