package filter

import (
	"fmt"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/autogen/app"
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

func newTitleFilter(filter *app.Filter) *TitleFilter {
	prefix, ok := filter.Props["prefix"]
	if !ok {
		prefix = "foo:"
	}
	return &TitleFilter{
		name:   "title",
		desc:   "This filter will prefix the title of the article with a given value.",
		tags:   filter.Tags,
		prefix: fmt.Sprintf("%v", prefix),
	}
}
