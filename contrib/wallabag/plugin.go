package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

var spec = model.Spec{
	Name: "wallabag",
	Desc: "Send new articles to a Wallabag instance.",
	PropsSpec: []model.PropSpec{
		{
			Name: "url",
			Desc: "Wallabag instance base URL",
			Type: model.Text,
		},
		{
			Name: "username",
			Desc: "Username",
			Type: model.Text,
		},
		{
			Name: "password",
			Desc: "Password",
			Type: model.Password,
		},
		{
			Name: "clientId",
			Desc: "Client ID",
			Type: model.Text,
		},
		{
			Name: "clientSecret",
			Desc: "Client secret",
			Type: model.Password,
		},
		{
			Name: "includeContent",
			Desc: "Include 'content' in API request",
			Type: model.Select,
			Options: map[string]string{
				"true":  "Yes",
				"false": "No (wallabag will fetch URL itself)",
			},
		},
	},
}

// WallabagOutputPlugin implements an Output Plugin for Wallabag
type WallabagOutputPlugin struct{}

// Spec returns plugin spec
func (p *WallabagOutputPlugin) Spec() model.Spec {
	return spec
}

// Build creates Twitter output provider instance
func (p *WallabagOutputPlugin) Build(def *model.OutputDef) (model.Output, error) {
	baseURL := def.Props.Get("url")
	if baseURL == "" {
		return nil, fmt.Errorf("missing url property")
	}
	username := def.Props.Get("username")
	if username == "" {
		return nil, fmt.Errorf("missing username property")
	}
	password := def.Props.Get("password")
	if password == "" {
		return nil, fmt.Errorf("missing password property")
	}
	clientID := def.Props.Get("clientId")
	if clientID == "" {
		return nil, fmt.Errorf("missing client ID property")
	}
	clientSecret := def.Props.Get("clientSecret")
	if clientSecret == "" {
		return nil, fmt.Errorf("missing client secret property")
	}

	includeContent := def.Props.Get("includeContent")
	if includeContent == "" {
		return nil, fmt.Errorf("missing includeContent property")
	}

	definition := *def
	definition.Spec = spec

	values := url.Values{}
	values.Add("username", username)
	values.Add("password", password)
	values.Add("grant_type", "password")
	oauthConf := &clientcredentials.Config{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		EndpointParams: values,
		TokenURL:       baseURL + "/oauth/v2/token",
	}

	return &WallabagOutputProvider{
		definition:     definition,
		baseURL:        baseURL,
		oauthConf:      oauthConf,
		includeContent: includeContent == "true",
	}, nil
}

// WallabagOutputProvider implements an output provider to send articles to Wallabag
type WallabagOutputProvider struct {
	definition     model.OutputDef
	baseURL        string
	oauthConf      *clientcredentials.Config
	includeContent bool
}

// Send sends an article to wallabag
func (op *WallabagOutputProvider) Send(article *model.Article) (bool, error) {
	client := op.oauthConf.Client(context.Background())

	values := url.Values{}
	values.Add("url", article.Link)
	values.Add("title", article.Title)
	values.Add("tags", strings.Join(article.Tags, ","))
	t := article.RefDate()
	if t != nil {
		values.Add("published_at", t.Format(time.RFC3339))
	}

	if op.includeContent {
		values.Add("content", article.Content)
	}

	resp, err := client.PostForm(op.baseURL+"/api/entries.json", values)
	if err != nil {
		atomic.AddUint32(&op.definition.NbError, 1)
		return false, err
	}
	defer resp.Body.Close()

	atomic.AddUint32(&op.definition.NbSuccess, 1)
	return true, err
}

// GetDef returns the plugin definition
func (op *WallabagOutputProvider) GetDef() model.OutputDef {
	return op.definition
}

// GetPluginSpec returns the plugin spec
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: spec,
		Type: model.OutputPluginType,
	}
}

// GetOutputPlugin returns the output plugin
func GetOutputPlugin() (op model.OutputPlugin, err error) {
	return &WallabagOutputPlugin{}, nil
}
