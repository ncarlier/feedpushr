package filter

import (
	"github.com/ncarlier/feedpushr/pkg/model"
)

type FooFilter struct {
	name string
}

func (f *FooFilter) DoFilter(article *model.Article) error {
	article.Title = "feedpushr: " + article.Title
	return nil
}

func newFooFilter() *FooFilter {
	return &FooFilter{
		name: "foo",
	}
}
