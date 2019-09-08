package main

import (
	"strconv"
	"sync/atomic"

	"github.com/k3a/html2text"
	"github.com/ncarlier/feedpushr/pkg/model"
)

//go:generate go run gen.go

var spec = model.Spec{
	Name: "rake",
	Desc: "Extract keywords of an article by using Rapid Automatic Keyword Extraction algorithm.",
	PropsSpec: []model.PropSpec{
		{
			Name: "minCharLength",
			Desc: "Minimum character length (default: 4)",
			Type: model.Number,
		},
		{
			Name: "maxWordsLength",
			Desc: "Maximum words length (default: 3)",
			Type: model.Number,
		},
		{
			Name: "minKeywordFrequency",
			Desc: "Minimum keyword frequency (default: 4)",
			Type: model.Number,
		},
	},
}

func safeAtoi(val string, fallback int) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return i
}

// RakeFilterPlugin is the RAKE filter plugin
type RakeFilterPlugin struct{}

// Spec returns plugin spec
func (p *RakeFilterPlugin) Spec() model.Spec {
	return spec
}

// Build creates RAKE filter
func (p *RakeFilterPlugin) Build(def *model.FilterDef) (model.Filter, error) {
	val := def.Props.Get("minCharLength")
	minCharLength := safeAtoi(val, 4)
	val = def.Props.Get("maxWordsLength")
	maxWordsLength := safeAtoi(val, 3)
	val = def.Props.Get("minKeywordFrequency")
	minKeywordFrequency := safeAtoi(val, 4)
	rake := NewRake("", minCharLength, maxWordsLength, minKeywordFrequency)
	rake.SetStopWords(stopWords)
	return &RakeFilter{
		id:      def.ID,
		alias:   def.Alias,
		spec:    spec,
		tags:    def.Tags,
		enabled: def.Enabled,
		rake:    rake,
	}, nil
}

// RakeFilter filter articles by adding extracted keywords
type RakeFilter struct {
	id        int
	alias     string
	spec      model.Spec
	tags      []string
	enabled   bool
	nbError   uint64
	nbSuccess uint64
	rake      *Rake
}

// DoFilter applies filter on the article
func (f *RakeFilter) DoFilter(article *model.Article) error {
	plain := html2text.HTML2Text(article.Content)
	if plain == "" {
		plain = article.Description
	}
	article.Meta["KeywordScore"] = f.rake.Run(plain)
	atomic.AddUint64(&f.nbSuccess, 1)
	return nil
}

// GetDef return output definition
func (f *RakeFilter) GetDef() model.FilterDef {
	result := model.FilterDef{
		ID:      f.id,
		Alias:   f.alias,
		Spec:    f.spec,
		Tags:    f.tags,
		Enabled: f.enabled,
	}
	result.Props = map[string]interface{}{
		"minCharLength":       f.rake.minCharLength,
		"maxWordsLength":      f.rake.maxWordsLength,
		"minKeywordFrequency": f.rake.minKeywordFrequency,
		"nbError":             f.nbError,
		"nbSuccess":           f.nbSuccess,
	}
	return result
}

// GetPluginSpec return plugin informations
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: spec,
		Type: model.FILTER_PLUGIN,
	}
}

// GetFilterPlugin returns filter plugin
func GetFilterPlugin() (op model.FilterPlugin, err error) {
	return &RakeFilterPlugin{}, nil
}
