package filter

import (
	"github.com/ncarlier/feedpushr/pkg/model"
)

// FooFilter is a foo filter
type FooFilter struct {
	name string
	desc string
}

// DoFilter applies filter on the article
func (f *FooFilter) DoFilter(article *model.Article) error {
	article.Title = "feedpushr: " + article.Title
	return nil
}

// GetSpec return filter specifications
func (f *FooFilter) GetSpec() model.FilterSpec {
	result := model.FilterSpec{
		Name: f.name,
		Desc: f.desc,
	}
	return result
}

func newFooFilter() *FooFilter {
	return &FooFilter{
		name: "foo",
		desc: "This filter will prefix the title of the article with the name of the application.",
	}
}
