package filter

import (
	"sync/atomic"

	"github.com/ncarlier/feedpushr/pkg/model"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/svg"
)

var minifySpec = model.Spec{
	Name:      "minify",
	Desc:      "This filter will minify articles HTML content.",
	PropsSpec: []model.PropSpec{},
}

// MinifyFilter is a filter that minify HTML content
type MinifyFilter struct {
	id        int
	spec      model.Spec
	tags      []string
	nbSuccess uint64
	nbError   uint64
	enabled   bool
	minifier  *minify.M
}

// DoFilter applies filter on the article
func (f *MinifyFilter) DoFilter(article *model.Article) error {
	if article.Content != "" {
		content, err := f.minifier.String("text/html", article.Content)
		if err != nil {
			atomic.AddUint64(&f.nbError, 1)
			return err
		}
		article.Content = content
	}
	if article.Text != "" {
		desc, err := f.minifier.String("text/html", article.Text)
		if err != nil {
			atomic.AddUint64(&f.nbError, 1)
			return err
		}
		article.Text = desc
	}

	atomic.AddUint64(&f.nbSuccess, 1)
	return nil
}

// GetDef return filter definition
func (f *MinifyFilter) GetDef() model.FilterDef {
	result := model.FilterDef{
		ID:      f.id,
		Tags:    f.tags,
		Spec:    f.spec,
		Enabled: f.enabled,
	}

	result.Props = map[string]interface{}{
		"nbSuccess": f.nbSuccess,
		"nbError":   f.nbError,
	}

	return result
}

func newMinifyFilter(filter *model.FilterDef) *MinifyFilter {
	minifier := minify.New()
	minifier.AddFunc("text/css", css.Minify)
	minifier.AddFunc("text/html", html.Minify)
	minifier.AddFunc("image/svg+xml", svg.Minify)
	return &MinifyFilter{
		id:       filter.ID,
		spec:     minifySpec,
		tags:     filter.Tags,
		minifier: minifier,
		enabled:  filter.Enabled,
	}
}
