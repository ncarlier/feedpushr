package main

import (
	"fmt"
	"net/url"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v3/pkg/format"
	"github.com/ncarlier/feedpushr/v3/pkg/format/fn"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
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
func (p *MastodonOutputPlugin) Build(def *model.OutputDef) (model.Output, error) {
	// Default format
	if frmt, ok := def.Props["format"]; !ok || frmt == "" {
		def.Props["format"] = "{{.Title}}\n{{.Link}}"
	}
	formatter, err := format.NewOutputFormatter(def)
	if err != nil {
		return nil, err
	}
	u := def.Props.Get("url")
	if u == "" {
		return nil, fmt.Errorf("missing URL property")
	}
	_url, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, fmt.Errorf("invalid URL property: %s", err.Error())
	}
	_url.Path = "/api/v1/statuses"
	accessToken := def.Props.Get("token")
	if accessToken == "" {
		return nil, fmt.Errorf("missing access token property")
	}
	visibility := def.Props.Get("visibility")
	if _, exists := tootVisibilities[visibility]; !exists {
		visibility = "public"
	}

	definition := *def
	definition.Spec = spec

	return &MastodonOutputProvider{
		definition:  definition,
		formatter:   formatter,
		targetURL:   _url.String(),
		accessToken: accessToken,
		visibility:  visibility,
	}, nil
}

// MastodonOutputProvider output provider to send articles to Mastodon
type MastodonOutputProvider struct {
	definition  model.OutputDef
	formatter   format.Formatter
	targetURL   string
	accessToken string
	visibility  string
}

// Send article to a Mastodon instance.
func (op *MastodonOutputProvider) Send(article *model.Article) (bool, error) {
	b, err := op.formatter.Format(article)
	if err != nil {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, err
	}
	toot := Toot{
		Status:     fn.Truncate(500, b.String()),
		Sensitive:  false,
		Visibility: op.visibility,
	}
	if err := sendToMastodon(toot, op.targetURL, op.accessToken); err != nil {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, err
	}
	atomic.AddUint64(&op.definition.NbSuccess, 1)
	return true, nil
}

// GetDef return output definition
func (op *MastodonOutputProvider) GetDef() model.OutputDef {
	return op.definition
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
