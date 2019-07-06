package filter

import (
	"sync/atomic"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/model"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/svg"
)

// MinifyFilter is a filter that minify HTML content
type MinifyFilter struct {
	name      string
	desc      string
	tags      []string
	nbSuccess uint64
	nbError   uint64
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
	if article.Description != "" {
		description, err := f.minifier.String("text/html", article.Description)
		if err != nil {
			atomic.AddUint64(&f.nbError, 1)
			return err
		}
		article.Description = description
	}

	atomic.AddUint64(&f.nbSuccess, 1)
	return nil
}

// GetSpec return filter specifications
func (f *MinifyFilter) GetSpec() model.FilterSpec {
	result := model.FilterSpec{
		Name: f.name,
		Desc: f.desc,
		Tags: f.tags,
	}
	result.Props = map[string]interface{}{
		"nbSuccess": f.nbSuccess,
		"nbError":   f.nbError,
	}

	return result
}

func newMinifyFilter(filter *app.Filter) *MinifyFilter {
	minifier := minify.New()
	minifier.AddFunc("text/css", css.Minify)
	minifier.AddFunc("text/html", html.Minify)
	minifier.AddFunc("image/svg+xml", svg.Minify)
	return &MinifyFilter{
		name:     "minify",
		desc:     "This filter will minify articles HTML content.",
		tags:     filter.Tags,
		minifier: minifier,
	}
}
