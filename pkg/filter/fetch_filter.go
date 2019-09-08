package filter

import (
	"context"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/pkg/model"
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
	tags      []string
	nbError   uint64
	nbSuccess uint64
	enabled   bool
}

// DoFilter applies filter on the article
func (f *FetchFilter) DoFilter(article *model.Article) error {
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
		ID:      f.id,
		Alias:   f.alias,
		Tags:    f.tags,
		Spec:    f.spec,
		Enabled: f.enabled,
	}

	result.Props = map[string]interface{}{
		"nbError":   f.nbError,
		"nbSuccess": f.nbSuccess,
	}
	return result
}

func newFetchFilter(filter *model.FilterDef) *FetchFilter {
	return &FetchFilter{
		id:      filter.ID,
		alias:   filter.Alias,
		spec:    fetchSpec,
		tags:    filter.Tags,
		enabled: filter.Enabled,
	}
}
