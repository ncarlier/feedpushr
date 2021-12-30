package plugins

import (
	"fmt"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v3/pkg/expr"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

var titleSpec = model.Spec{
	Name: "title",
	Desc: "This filter will prefix the title of the article with a given value.",
	PropsSpec: []model.PropSpec{
		{
			Name: "prefix",
			Desc: "Prefix to add to the article title",
			Type: model.Text,
		},
	},
}

// TitleFilterPlugin is the RAKE filter plugin
type TitleFilterPlugin struct{}

// Spec returns plugin spec
func (p *TitleFilterPlugin) Spec() model.Spec {
	return titleSpec
}

// Build creates RAKE filter
func (p *TitleFilterPlugin) Build(def *model.FilterDef) (model.Filter, error) {
	condition, err := expr.NewConditionalExpression(def.Condition)
	if err != nil {
		return nil, err
	}
	definition := *def
	definition.Spec = titleSpec
	prefix, ok := definition.Props["prefix"]
	if !ok {
		prefix = "feedpushr:"
		definition.Props["prefix"] = prefix
	}
	return &TitleFilter{
		definition: definition,
		condition:  condition,
		prefix:     fmt.Sprintf("%v", prefix),
	}, nil
}

// TitleFilter is a foo filter
type TitleFilter struct {
	definition model.FilterDef
	condition  *expr.ConditionalExpression
	prefix     string
}

// DoFilter applies filter on the article
func (f *TitleFilter) DoFilter(article *model.Article) (bool, error) {
	article.Title = f.prefix + " " + article.Title
	atomic.AddUint32(&f.definition.NbSuccess, 1)
	return true, nil
}

// Match test if article matches filter condition
func (f *TitleFilter) Match(article *model.Article) bool {
	return f.condition.Match(article)
}

// GetDef return filter definition
func (f *TitleFilter) GetDef() model.FilterDef {
	return f.definition
}
