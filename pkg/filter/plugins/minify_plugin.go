package plugins

import (
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v3/pkg/expr"
	"github.com/ncarlier/feedpushr/v3/pkg/model"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/svg"
)

var minifySpec = model.Spec{
	Name:      "minify",
	Desc:      "This filter will minify articles HTML content.",
	PropsSpec: []model.PropSpec{},
}

// MinifyFilterPlugin is the minification filter plugin
type MinifyFilterPlugin struct{}

// Spec returns plugin spec
func (p *MinifyFilterPlugin) Spec() model.Spec {
	return minifySpec
}

// Build creates minification filter
func (p *MinifyFilterPlugin) Build(def *model.FilterDef) (model.Filter, error) {
	condition, err := expr.NewConditionalExpression(def.Condition)
	if err != nil {
		return nil, err
	}
	definition := *def
	definition.Spec = minifySpec
	minifier := minify.New()
	minifier.AddFunc("text/css", css.Minify)
	minifier.AddFunc("text/html", html.Minify)
	minifier.AddFunc("image/svg+xml", svg.Minify)
	return &MinifyFilter{
		definition: definition,
		condition:  condition,
		minifier:   minifier,
	}, nil
}

// MinifyFilter is a filter that minify HTML content
type MinifyFilter struct {
	definition model.FilterDef
	condition  *expr.ConditionalExpression
	minifier   *minify.M
}

// DoFilter applies filter on the article
func (f *MinifyFilter) DoFilter(article *model.Article) (bool, error) {
	if article.Content != "" {
		content, err := f.minifier.String("text/html", article.Content)
		if err != nil {
			atomic.AddUint32(&f.definition.NbError, 1)
			return false, err
		}
		article.Content = content
	}
	if article.Text != "" {
		desc, err := f.minifier.String("text/html", article.Text)
		if err != nil {
			atomic.AddUint32(&f.definition.NbError, 1)
			return false, err
		}
		article.Text = desc
	}

	atomic.AddUint32(&f.definition.NbSuccess, 1)
	return true, nil
}

// Match test if article matches filter condition
func (f *MinifyFilter) Match(article *model.Article) bool {
	return f.condition.Match(article)
}

// GetDef return filter definition
func (f *MinifyFilter) GetDef() model.FilterDef {
	return f.definition
}
