package filter

import (
	"github.com/ncarlier/feedpushr/pkg/model"
)

// FooFilter is a foo filter
type FooFilter struct {
	name string
}

// DoFilter applies filter on the article
func (f *FooFilter) DoFilter(article *model.Article) error {
	article.Title = "feedpushr: " + article.Title
	return nil
}

func newFooFilter() *FooFilter {
	return &FooFilter{
		name: "foo",
	}
}
