package plugins

import (
	"context"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v3/pkg/expr"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/readflow/pkg/readability"
)

var fetchSpec = model.Spec{
	Name: "fetch",
	Desc: `
This filter will attempt to extract the content of the article from the source URL.
If succeeded, following metadata are added to the article:

- originalContent: Initial article content (before fetching)
- text: Article excerpt
- image: Article main illustration
`,
	PropsSpec: []model.PropSpec{},
}

// FetchFilterPlugin is the fetch filter plugin
type FetchFilterPlugin struct{}

// Spec returns plugin spec
func (p *FetchFilterPlugin) Spec() model.Spec {
	return fetchSpec
}

// Build creates fetch filter
func (p *FetchFilterPlugin) Build(def *model.FilterDef) (model.Filter, error) {
	condition, err := expr.NewConditionalExpression(def.Condition)
	if err != nil {
		return nil, err
	}
	definition := *def
	definition.Spec = fetchSpec
	return &FetchFilter{
		definition: definition,
		condition:  condition,
	}, nil
}

// FetchFilter is a filter that try to fetch the original article content
type FetchFilter struct {
	definition model.FilterDef
	condition  *expr.ConditionalExpression
}

// DoFilter applies filter on the article
func (f *FetchFilter) DoFilter(article *model.Article) (bool, error) {
	art, err := readability.FetchArticle(context.Background(), article.Link)
	if err != nil && art == nil {
		atomic.AddUint64(&f.definition.NbError, 1)
		return false, err
	}
	article.Title = art.Title
	if art.HTML != nil {
		article.Meta["originalContent"] = article.Content
		article.Content = *art.HTML
	}
	if art.Text != nil {
		article.Meta["text"] = *art.Text
	}
	if art.Image != nil {
		article.Meta["image"] = *art.Image
	}
	// article.Meta["length"] = art.Length
	// article.Meta["sitename"] = art.SiteName
	atomic.AddUint64(&f.definition.NbSuccess, 1)
	return true, nil
}

// Match test if article matches filter condition
func (f *FetchFilter) Match(article *model.Article) bool {
	return f.condition.Match(article)
}

// GetDef return filter definition
func (f *FetchFilter) GetDef() model.FilterDef {
	return f.definition
}
