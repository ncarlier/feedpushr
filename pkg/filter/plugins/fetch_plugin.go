package plugins

import (
	"context"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v3/pkg/expr"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/readflow/pkg/scraper"
)

var fetchSpec = model.Spec{
	Name: "fetch",
	Desc: `
This filter will attempt to extract the content of the article from the source URL.
If succeeded, following metadata are added (if found) to the article:

- originalContent: Initial article content (before fetching)
- image: Article main illustration
- excerpt: Article excerpt
- length: Article length
- sitename: Website name
- favicon: Website favicon
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
		scraper:    scraper.NewInternalWebScraper(),
	}, nil
}

// FetchFilter is a filter that try to fetch the original article content
type FetchFilter struct {
	definition model.FilterDef
	condition  *expr.ConditionalExpression
	scraper    scraper.WebScraper
}

// DoFilter applies filter on the article
func (f *FetchFilter) DoFilter(article *model.Article) (bool, error) {
	webpage, err := f.scraper.Scrap(context.Background(), article.Link)
	if err != nil && webpage == nil {
		atomic.AddUint64(&f.definition.NbError, 1)
		return false, err
	}
	article.Title = webpage.Title
	if webpage.HTML != "" {
		article.Meta["originalContent"] = article.Content
		article.Content = webpage.HTML
	}
	if webpage.Text != "" {
		article.Text = webpage.Text
	}

	// Add meta...
	article.Meta["excerpt"] = webpage.Excerpt
	article.Meta["image"] = webpage.Image
	article.Meta["sitename"] = webpage.SiteName
	article.Meta["favicon"] = webpage.Favicon
	article.Meta["length"] = webpage.Length

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
