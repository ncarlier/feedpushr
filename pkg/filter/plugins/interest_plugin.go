package plugins

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/expr"
	"github.com/ncarlier/feedpushr/v3/pkg/model"

	"github.com/cloudflare/ahocorasick"
)

var interestSpec = model.Spec{
	Name: "interest",
	Desc: "This filter will keep articles that match a list of words of interest.",
	PropsSpec: []model.PropSpec{
		{
			Name: "wordlist",
			Desc: "Comma-sperated list of words of interest",
			Type: model.Textarea,
		},
	},
}

// InterestFilterPlugin is the Interest filter plugin
type InterestFilterPlugin struct{}

// Spec returns plugin spec
func (p *InterestFilterPlugin) Spec() model.Spec {
	return interestSpec
}

// Build creates interest filter
func (p *InterestFilterPlugin) Build(def *model.FilterDef) (model.Filter, error) {
	condition, err := expr.NewConditionalExpression(def.Condition)
	if err != nil {
		return nil, err
	}
	definition := *def
	definition.Spec = interestSpec

	wordlist, ok := definition.Props["wordlist"]
	if !ok {
		wordlist = "news"
		definition.Props["wordlist"] = wordlist
	}

	patterns := strings.Split(fmt.Sprintf("%v", wordlist), ",")
	for i := range patterns {
		patterns[i] = strings.TrimSpace(patterns[i])
	}
	matcher := ahocorasick.NewStringMatcher(patterns)

	return &InterestFilter{
		definition: definition,
		condition:  condition,
		matcher:    matcher,
	}, nil
}

// InterestFilter is a interest filter
type InterestFilter struct {
	definition model.FilterDef
	condition  *expr.ConditionalExpression
	matcher    *ahocorasick.Matcher
}

// DoFilter applies filter on the article
func (f *InterestFilter) DoFilter(article *model.Article) (bool, error) {
	if f.matcher.Contains([]byte(article.Title)) || f.matcher.Contains([]byte(article.Text)) {
		atomic.AddUint32(&f.definition.NbSuccess, 1)
		return true, nil
	}
	return true, common.ErrArticleShouldBeIgnored
}

// Interest test if article matches filter condition
func (f *InterestFilter) Match(article *model.Article) bool {
	return f.condition.Match(article)
}

// GetDef return filter definition
func (f *InterestFilter) GetDef() model.FilterDef {
	return f.definition
}
