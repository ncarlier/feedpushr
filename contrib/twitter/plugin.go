package main

import (
	"fmt"
	"net/url"
	"strings"
	"sync/atomic"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ncarlier/feedpushr/v2/pkg/format"
	"github.com/ncarlier/feedpushr/v2/pkg/format/fn"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

var spec = model.Spec{
	Name: "twitter",
	Desc: "Send new articles to a Twitter timeline.",
	PropsSpec: []model.PropSpec{
		{
			Name: "consumerKey",
			Desc: "Consumer key",
			Type: model.Text,
		},
		{
			Name: "consumerSecret",
			Desc: "Consumer secret",
			Type: model.Password,
		},
		{
			Name: "accessToken",
			Desc: "Access token",
			Type: model.Text,
		},
		{
			Name: "accessTokenSecret",
			Desc: "Access token secret",
			Type: model.Password,
		},
		{
			Name: "format",
			Desc: "Tweet format (default: `{{tweet .Title .Link}}`)",
			Type: model.Textarea,
		},
	},
}

// TwitterOutputPlugin is the Twitter output plugin
type TwitterOutputPlugin struct{}

// Spec returns plugin spec
func (p *TwitterOutputPlugin) Spec() model.Spec {
	return spec
}

// Build creates Twitter output provider instance
func (p *TwitterOutputPlugin) Build(def *model.OutputDef) (model.Output, error) {
	// Default format
	if frmt, ok := def.Props["format"]; !ok || frmt == "" {
		def.Props["format"] = "{{tweet .Title .Link}}"
	}
	formatter, err := format.NewOutputFormatter(def)
	if err != nil {
		return nil, err
	}
	consumerKey := def.Props.Get("consumerKey")
	if consumerKey == "" {
		return nil, fmt.Errorf("missing consumer key property")
	}
	consumerSecret := def.Props.Get("consumerSecret")
	if consumerSecret == "" {
		return nil, fmt.Errorf("missing consumer secret property")
	}
	accessToken := def.Props.Get("accessToken")
	if accessToken == "" {
		return nil, fmt.Errorf("missing access token property")
	}
	accessTokenSecret := def.Props.Get("accessTokenSecret")
	if accessTokenSecret == "" {
		return nil, fmt.Errorf("missing access token secret property")
	}
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	definition := *def
	definition.Spec = spec

	return &TwitterOutputProvider{
		definition:     definition,
		formatter:      formatter,
		api:            api,
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
	}, nil
}

// TwitterOutputProvider output provider to send articles to Twitter
type TwitterOutputProvider struct {
	definition     model.OutputDef
	formatter      format.Formatter
	consumerKey    string
	consumerSecret string
	api            *anaconda.TwitterApi
}

// Send sent an article as Tweet to a Twitter timeline
func (op *TwitterOutputProvider) Send(article *model.Article) (bool, error) {
	b, err := op.formatter.Format(article)
	if err != nil {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, err
	}
	tweet := fn.Truncate(270, b.String())
	v := url.Values{}
	_, err = op.api.PostTweet(tweet, v)
	if err != nil {
		// Ignore error due to duplicate status
		if strings.Contains(err.Error(), "\"code\":187") {
			return true, nil
		}
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, nil
	}
	atomic.AddUint64(&op.definition.NbSuccess, 1)
	return true, err
}

// GetDef return filter definition
func (op *TwitterOutputProvider) GetDef() model.OutputDef {
	return op.definition
}

// GetPluginSpec returns plugin spec
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: spec,
		Type: model.OutputPluginType,
	}
}

// GetOutputPlugin returns output plugin
func GetOutputPlugin() (op model.OutputPlugin, err error) {
	return &TwitterOutputPlugin{}, nil
}
