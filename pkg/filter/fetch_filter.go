package filter

import (
	"context"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v2/pkg/expr"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
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

// FetchFilter is a filter that try to fetch the original article content
type FetchFilter struct {
	id        int
	alias     string
	spec      model.Spec
	condition *expr.ConditionalExpression
	nbError   uint64
	nbSuccess uint64
	enabled   bool
}

// DoFilter applies filter on the article
func (f *FetchFilter) DoFilter(article *model.Article) error {
	if !f.enabled || !f.condition.Match(article) {
		// Ignore if disabled or if the article doesn't match the condition
		return nil
	}
	art, err := readability.FetchArticle(context.Background(), article.Link)
	if err != nil && art == nil {
		atomic.AddUint64(&f.nbError, 1)
		return err
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
	atomic.AddUint64(&f.nbSuccess, 1)
	return nil
}

// GetDef return filter definition
func (f *FetchFilter) GetDef() model.FilterDef {
	result := model.FilterDef{
		ID:        f.id,
		Alias:     f.alias,
		Spec:      f.spec,
		Condition: f.condition.String(),
		Enabled:   f.enabled,
	}

	result.Props = map[string]interface{}{
		"nbError":   f.nbError,
		"nbSuccess": f.nbSuccess,
	}
	return result
}

func newFetchFilter(filter *model.FilterDef) (*FetchFilter, error) {
	condition, err := expr.NewConditionalExpression(filter.Condition)
	if err != nil {
		return nil, err
	}
	return &FetchFilter{
		id:        filter.ID,
		alias:     filter.Alias,
		spec:      fetchSpec,
		condition: condition,
		enabled:   filter.Enabled,
	}, nil
}
