package filter

import (
	"fmt"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v2/pkg/expr"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

var titleSpec = model.Spec{
	Name: "title",
	Desc: "This filter will prefix the title of the article with a given value.",
	PropsSpec: []model.PropSpec{
		{
			Name: "prefix",
			Desc: "Prefix to add to the article title",
			Type: model.Text,
		},
	},
}

// TitleFilter is a foo filter
type TitleFilter struct {
	id        int
	alias     string
	spec      model.Spec
	condition *expr.ConditionalExpression
	prefix    string
	nbSuccess uint64
	enabled   bool
}

// DoFilter applies filter on the article
func (f *TitleFilter) DoFilter(article *model.Article) error {
	if !f.enabled || !f.condition.Match(article) {
		// Ignore if disabled or if the article doesn't match the condition
		return nil
	}
	article.Title = f.prefix + " " + article.Title
	atomic.AddUint64(&f.nbSuccess, 1)
	return nil
}

// GetDef return filter definition
func (f *TitleFilter) GetDef() model.FilterDef {
	result := model.FilterDef{
		ID:        f.id,
		Alias:     f.alias,
		Condition: f.condition.String(),
		Spec:      f.spec,
		Enabled:   f.enabled,
	}

	result.Props = map[string]interface{}{
		"prefix":    f.prefix,
		"nbSuccess": f.nbSuccess,
	}

	return result
}

func newTitleFilter(filter *model.FilterDef) (*TitleFilter, error) {
	condition, err := expr.NewConditionalExpression(filter.Condition)
	if err != nil {
		return nil, err
	}
	prefix, ok := filter.Props["prefix"]
	if !ok {
		prefix = "feedpushr:"
	}
	return &TitleFilter{
		id:        filter.ID,
		alias:     filter.Alias,
		spec:      titleSpec,
		condition: condition,
		prefix:    fmt.Sprintf("%v", prefix),
		enabled:   filter.Enabled,
	}, nil
}
