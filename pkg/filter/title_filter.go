package filter

import (
	"net/url"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/model"
)

// TitleFilter is a foo filter
type TitleFilter struct {
	name      string
	desc      string
	tags      []string
	prefix    string
	nbSuccess uint64
}

// DoFilter applies filter on the article
func (f *TitleFilter) DoFilter(article *model.Article) error {
	article.Title = f.prefix + " " + article.Title
	atomic.AddUint64(&f.nbSuccess, 1)
	return nil
}

// GetSpec return filter specifications
func (f *TitleFilter) GetSpec() model.FilterSpec {
	result := model.FilterSpec{
		Name: f.name,
		Desc: f.desc,
		Tags: f.tags,
	}
	result.Props = map[string]interface{}{
		"prefix":     f.prefix,
		"nbSsuccess": f.nbSuccess,
	}

	return result
}

func newTitleFilter(params url.Values, tags string) *TitleFilter {
	return &TitleFilter{
		name:   "title",
		desc:   "This filter will prefix the title of the article with a given value.",
		tags:   builder.GetFeedTags(&tags),
		prefix: params.Get("prefix"),
	}
}
