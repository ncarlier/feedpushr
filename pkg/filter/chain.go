package filter

import (
	"fmt"
	"net/url"

	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/feedpushr/pkg/plugin"
)

// Chain contains filter chain
type Chain struct {
	filters []model.Filter
}

// NewChainFilter creates a new filter chain
func NewChainFilter(filters []string, pr *plugin.Registry) (*Chain, error) {
	chain := &Chain{}

	for _, f := range filters {
		u, err := url.Parse(f)
		if err != nil {
			return nil, fmt.Errorf("invalid filter URL: %s", f)
		}
		switch u.Scheme {
		case "title":
			chain.filters = append(chain.filters, newTitleFilter(u.Query(), u.Fragment))
		case "fetch":
			chain.filters = append(chain.filters, newFetchFilter(u.Fragment))
		case "minify":
			chain.filters = append(chain.filters, newMinifyFilter(u.Query(), u.Fragment))
		default:
			// Try to load plugin regarding the name
			plug := pr.LookupFilterPlugin(u.Scheme)
			if plug == nil {
				return nil, fmt.Errorf("unsuported filter: %s", u.Scheme)
			}
			fp, err := plug.Build(u.Query(), u.Fragment)
			if err != nil {
				return nil, fmt.Errorf("unable to create filter: %v", err)
			}
			chain.filters = append(chain.filters, fp)
		}
	}

	return chain, nil
}

// Apply applies filter chain on an article
func (c *Chain) Apply(article *model.Article) error {
	for idx, filter := range c.filters {
		tags := filter.GetSpec().Tags
		if !match(tags, article.Tags) {
			// Ignore filters that do not match the article tags
			continue
		}
		err := filter.DoFilter(article)
		if err != nil {
			return fmt.Errorf("error while applying filter #%d: %v", idx, err)
		}
	}
	return nil
}

// GetSpec return specification of the chain filter
func (c *Chain) GetSpec() []model.FilterSpec {
	result := make([]model.FilterSpec, len(c.filters))
	for idx, filter := range c.filters {
		result[idx] = filter.GetSpec()
	}
	return result
}

func match(a, b []string) bool {
	// A filter with no tags match all articles
	if len(a) == 0 {
		return true
	}
	bSet := make(map[string]struct{}, len(b))
	for _, s := range b {
		bSet[s] = struct{}{}
	}

	for _, tag := range a {
		if _, ok := bSet[tag]; !ok {
			return false
		}
	}
	return true
}
