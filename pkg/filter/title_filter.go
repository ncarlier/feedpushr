package filter

import (
	"fmt"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/pkg/model"
)

var titleSpec = model.Spec{
	Name: "title",
	Desc: "This filter will prefix the title of the article with a given value.",
	PropsSpec: []model.PropSpec{
		{
			Name: "prefix",
			Desc: "Prefix to add to the article title",
			Type: "string",
		},
	},
}

// TitleFilter is a foo filter
type TitleFilter struct {
	id        int
	spec      model.Spec
	tags      []string
	prefix    string
	nbSuccess uint64
	enabled   bool
}

// DoFilter applies filter on the article
func (f *TitleFilter) DoFilter(article *model.Article) error {
	article.Title = f.prefix + " " + article.Title
	atomic.AddUint64(&f.nbSuccess, 1)
	return nil
}

// GetDef return filter definition
func (f *TitleFilter) GetDef() model.FilterDef {
	result := model.FilterDef{
		ID:      f.id,
		Tags:    f.tags,
		Spec:    f.spec,
		Enabled: f.enabled,
	}

	result.Props = map[string]interface{}{
		"prefix":    f.prefix,
		"nbSuccess": f.nbSuccess,
	}

	return result
}

func newTitleFilter(filter *model.FilterDef) *TitleFilter {
	prefix, ok := filter.Props["prefix"]
	if !ok {
		prefix = "foo:"
	}
	return &TitleFilter{
		id:      filter.ID,
		spec:    titleSpec,
		tags:    filter.Tags,
		prefix:  fmt.Sprintf("%v", prefix),
		enabled: filter.Enabled,
	}
}
