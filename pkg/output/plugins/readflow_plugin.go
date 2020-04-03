package plugins

import (
	"fmt"
	"net/url"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

var readflowSpec = model.Spec{
	Name: "readflow",
	Desc: "Send new articles a readflow instance.",
	PropsSpec: []model.PropSpec{
		{
			Name: "url",
			Desc: "Target URL",
			Type: model.Text,
		},
		{
			Name: "apiKey",
			Desc: "API KEY",
			Type: model.Password,
		},
	},
}

// ReadflowOutputPlugin is the Readflow output plugin
type ReadflowOutputPlugin struct{}

// Spec returns plugin spec
func (p *ReadflowOutputPlugin) Spec() model.Spec {
	return readflowSpec
}

// Build creates Readflow output provider instance
func (p *ReadflowOutputPlugin) Build(def *model.OutputDef) (model.Output, error) {
	u := def.Props.Get("url")
	if u == "" {
		u = "https://api.readflow.app"
	}
	targetURL, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, fmt.Errorf("invalid URL property: %s", err.Error())
	}
	apiKey := def.Props.Get("apiKey")
	if apiKey == "" {
		return nil, fmt.Errorf("missing API key property")
	}
	definition := *def
	definition.Spec = readflowSpec
	definition.Props["url"] = targetURL.String()
	return &ReadflowOutputProvider{
		definition: definition,
		targetURL:  targetURL.String(),
		apiKey:     apiKey,
	}, nil
}

// ReadflowOutputProvider output provider to send articles to Readflow
type ReadflowOutputProvider struct {
	definition model.OutputDef
	targetURL  string
	apiKey     string
}

// Send article to a Readflow instance.
func (op *ReadflowOutputProvider) Send(article *model.Article) (bool, error) {
	nb, err := sendToReadflow(op.targetURL, op.apiKey, article)
	if err != nil {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, err
	}
	atomic.AddUint64(&op.definition.NbSuccess, uint64(nb))
	return true, nil
}

// GetDef return output definition
func (op *ReadflowOutputProvider) GetDef() model.OutputDef {
	return op.definition
}

// GetPluginSpec return plugin informations
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: readflowSpec,
		Type: model.OutputPluginType,
	}
}

// GetOutputPlugin return output plugin instance
func GetOutputPlugin() (op model.OutputPlugin, err error) {
	return &ReadflowOutputPlugin{}, nil
}
