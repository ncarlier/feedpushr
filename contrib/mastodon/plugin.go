package main

import (
	"fmt"
	"net/url"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/pkg/model"
)

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
			Name: "visibility",
			Desc: "Toot visibiliy (public, unlisted, private, direct)",
			Type: model.Text,
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
func (p *MastodonOutputPlugin) Build(output *model.OutputDef) (model.OutputProvider, error) {
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
	if visibility == "" {
		visibility = "public"
	}
	return &MastodonOutputProvider{
		id:          output.ID,
		spec:        spec,
		tags:        output.Tags,
		enabled:     output.Enabled,
		targetURL:   _url.String(),
		accessToken: accessToken,
		visibility:  visibility,
	}, nil
}

// MastodonOutputProvider output provider to send articles to Mastodon
type MastodonOutputProvider struct {
	id          int
	spec        model.Spec
	tags        []string
	enabled     bool
	nbError     uint64
	nbSuccess   uint64
	targetURL   string
	accessToken string
	visibility  string
}

// Send article to a Mastodon instance.
func (op *MastodonOutputProvider) Send(article *model.Article) error {
	toot := Toot{
		Status:     fmt.Sprintf("%s\n%s", article.Title, article.Link),
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
		ID:      op.id,
		Spec:    op.spec,
		Tags:    op.tags,
		Enabled: op.enabled,
	}
	result.Props = map[string]interface{}{
		"url":         op.targetURL,
		"accessToken": op.accessToken,
		"visibility":  op.visibility,
		"nbError":     op.nbError,
		"nbSuccess":   op.nbSuccess,
	}
	return result
}

// GetPluginSpec return plugin informations
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: spec,
		Type: model.OUTPUT_PLUGIN,
	}
}

// GetOutputPlugin returns output plugin
func GetOutputPlugin() (op model.OutputPlugin, err error) {
	return &MastodonOutputPlugin{}, nil
}