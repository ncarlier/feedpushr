package main

import (
	"fmt"
	"net/url"
	"sync/atomic"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ncarlier/feedpushr/pkg/model"
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
	},
}

// TwitterOutputPlugin is the Twitter output plugin
type TwitterOutputPlugin struct{}

// Spec returns plugin spec
func (p *TwitterOutputPlugin) Spec() model.Spec {
	return spec
}

// Build creates Twitter output provider instance
func (p *TwitterOutputPlugin) Build(output *model.OutputDef) (model.OutputProvider, error) {
	consumerKey := output.Props.Get("consumerKey")
	if consumerKey == "" {
		return nil, fmt.Errorf("missing consumer key property")
	}
	consumerSecret := output.Props.Get("consumerSecret")
	if consumerSecret == "" {
		return nil, fmt.Errorf("missing consumer secret property")
	}
	accessToken := output.Props.Get("accessToken")
	if accessToken == "" {
		return nil, fmt.Errorf("missing access token property")
	}
	accessTokenSecret := output.Props.Get("accessTokenSecret")
	if accessTokenSecret == "" {
		return nil, fmt.Errorf("missing access token secret property")
	}
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	return &TwitterOutputProvider{
		id:      output.ID,
		spec:    spec,
		tags:    output.Tags,
		enabled: output.Enabled,
		api:     api,
	}, nil
}

// TwitterOutputProvider output provider to send articles to Twitter
type TwitterOutputProvider struct {
	id        int
	spec      model.Spec
	tags      []string
	enabled   bool
	nbError   uint64
	nbSuccess uint64
	// TODO add credentials
	api *anaconda.TwitterApi
}

// Send sent an article as Tweet to a Twitter timeline
func (op *TwitterOutputProvider) Send(article *model.Article) error {
	tweet := fmt.Sprintf("%s\n%s", article.Title, article.Link)
	r := []rune(tweet)
	if len(r) > 280 {
		nbCharToTruncate := len(r) - 280
		title := []rune(article.Title)
		end := len(title) - nbCharToTruncate
		tweet = fmt.Sprintf("%s\n%s", string(title[:end]), article.Link)
	}
	v := url.Values{}
	_, err := op.api.PostTweet(tweet, v)
	if err != nil {
		atomic.AddUint64(&op.nbError, 1)
	} else {
		atomic.AddUint64(&op.nbSuccess, 1)
	}
	return err
}

// GetDef return filter definition
func (op *TwitterOutputProvider) GetDef() model.OutputDef {
	result := model.OutputDef{
		Spec: op.spec,
		Tags: op.tags,
	}
	result.Props = map[string]interface{}{
		"accessToken":       op.api.Credentials.Token,
		"accessTokenSecret": model.MaskSecret(op.api.Credentials.Secret),
		"nbError":           op.nbError,
		"nbSuccess":         op.nbSuccess,
	}
	return result
}

// GetPluginSpec returns plugin spec
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: spec,
		Type: model.OUTPUT_PLUGIN,
	}
}

// GetOutputPlugin returns output plugin
func GetOutputPlugin() (op model.OutputPlugin, err error) {
	return &TwitterOutputPlugin{}, nil
}
