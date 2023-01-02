package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v3/pkg/expr"
	"github.com/ncarlier/feedpushr/v3/pkg/format"
	httpc "github.com/ncarlier/feedpushr/v3/pkg/http"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

var httpSpec = model.Spec{
	Name: "http",
	Desc: `
This filter will send the article as JSON object to a HTTP endpoint (POST).

HTTP endpoint must return same [JSON structure](https://github.com/ncarlier/feedpushr#output-format).
If succeeded, the response is merged with the current article.
`,
	PropsSpec: []model.PropSpec{
		{
			Name: "url",
			Desc: "Target URL",
			Type: model.URL,
		},
	},
}

// HTTPFilterPlugin is the http filter plugin
type HTTPFilterPlugin struct{}

// Spec returns plugin spec
func (p *HTTPFilterPlugin) Spec() model.Spec {
	return httpSpec
}

// Build creates http filter
func (p *HTTPFilterPlugin) Build(def *model.FilterDef) (model.Filter, error) {
	condition, err := expr.NewConditionalExpression(def.Condition)
	if err != nil {
		return nil, err
	}

	u, ok := def.Props["url"]
	if !ok {
		return nil, fmt.Errorf("missing URL property")
	}
	targetURL, err := url.ParseRequestURI(fmt.Sprintf("%v", u))
	if err != nil {
		return nil, fmt.Errorf("invalid URL property: %s", err.Error())
	}

	definition := *def
	definition.Spec = httpSpec
	definition.Props["url"] = targetURL.String()

	return &HTTPFilter{
		definition: definition,
		condition:  condition,
		targetURL:  targetURL.String(),
		formatter:  format.NewJSONFormatter(),
	}, nil
}

// HTTPFilter is a filter that try to http the original article content
type HTTPFilter struct {
	definition model.FilterDef
	condition  *expr.ConditionalExpression
	targetURL  string
	formatter  format.Formatter
}

// DoFilter applies filter on the article
func (f *HTTPFilter) DoFilter(article *model.Article) (bool, error) {
	b, err := f.formatter.Format(article)
	if err != nil {
		atomic.AddUint32(&f.definition.NbError, 1)
		return false, err
	}

	req, err := http.NewRequest("POST", f.targetURL, b)
	if err != nil {
		atomic.AddUint32(&f.definition.NbError, 1)
		return false, err
	}
	req.Header.Set("User-Agent", httpc.UserAgent)
	req.Header.Set("Content-Type", httpc.ContentTypeJSON)
	resp, err := httpc.DefaultClient.Do(req)
	if err != nil {
		atomic.AddUint32(&f.definition.NbError, 1)
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		atomic.AddUint32(&f.definition.NbError, 1)
		return false, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	var returnedArticle model.Article
	if err := json.NewDecoder(resp.Body).Decode(&returnedArticle); err != nil {
		atomic.AddUint32(&f.definition.NbError, 1)
		return false, fmt.Errorf("invalid JSON payload: %s", err.Error())
	}

	// Merge article with the response
	article.Merge(returnedArticle)

	atomic.AddUint32(&f.definition.NbSuccess, 1)
	return true, nil
}

// Match test if article matches filter condition
func (f *HTTPFilter) Match(article *model.Article) bool {
	return f.condition.Match(article)
}

// GetDef return filter definition
func (f *HTTPFilter) GetDef() model.FilterDef {
	return f.definition
}
