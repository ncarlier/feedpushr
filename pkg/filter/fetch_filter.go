package filter

import (
	"context"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/readflow/pkg/readability"
)

// FetchFilter is a filter that try to fetch the original article content
type FetchFilter struct {
	name      string
	desc      string
	tags      []string
	nbError   uint64
	nbSuccess uint64
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

// GetSpec return filter specifications
func (f *FetchFilter) GetSpec() model.FilterSpec {
	result := model.FilterSpec{
		Name: f.name,
		Desc: f.desc,
		Tags: f.tags,
	}
	result.Props = map[string]interface{}{
		"nbError":    f.nbError,
		"nbSsuccess": f.nbSuccess,
	}
	return result
}

const fetchDescription = `
This filter will attempt to extract the content of the article from the source URL.
If succeeded, following metadata are added to the article:

- originalContent: Initial article content (before fetching)
- text: Article excerpt
- image: Article main illustration
`

func newFetchFilter(tags []string) *FetchFilter {
	return &FetchFilter{
		name: "fetch",
		desc: fetchDescription,
		tags: tags,
	}
}
