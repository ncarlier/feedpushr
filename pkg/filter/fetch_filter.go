package filter

import (
	"time"

	"github.com/RadhiFadlillah/go-readability"
	"github.com/ncarlier/feedpushr/pkg/model"
)

// FetchFilter is a filter that try to fetch the original article content
type FetchFilter struct {
	name string
}

// DoFilter applies filter on the article
func (f *FetchFilter) DoFilter(article *model.Article) error {
	art, err := readability.Parse(article.Link, 5*time.Second)
	if err != nil {
		return err
	}
	article.Meta["RawContent"] = article.Content
	article.Content = art.Content
	article.Meta["MinReadTime"] = art.Meta.MinReadTime
	article.Meta["MaxReadTime"] = art.Meta.MaxReadTime
	return nil
}

func newFetchFilter() *FetchFilter {
	return &FetchFilter{
		name: "fetch",
	}
}
