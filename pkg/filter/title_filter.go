package filter

import (
	"net/url"

	"github.com/ncarlier/feedpushr/pkg/model"
)

// TitleFilter is a foo filter
type TitleFilter struct {
	name   string
	desc   string
	prefix string
}

// DoFilter applies filter on the article
func (f *TitleFilter) DoFilter(article *model.Article) error {
	article.Title = f.prefix + " " + article.Title
	return nil
}

// GetSpec return filter specifications
func (f *TitleFilter) GetSpec() model.FilterSpec {
	result := model.FilterSpec{
		Name: f.name,
		Desc: f.desc,
	}
	result.Props = map[string]interface{}{
		"prefix": f.prefix,
	}

	return result
}

func newTitleFilter(params url.Values) *TitleFilter {
	return &TitleFilter{
		name:   "title",
		desc:   "This filter will prefix the title of the article with a given value.",
		prefix: params.Get("prefix"),
	}
}
