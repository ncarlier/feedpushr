package filter

import (
	"fmt"

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

	for _, name := range filters {
		switch name {
		case "foo":
			chain.filters = append(chain.filters, newFooFilter())
		case "fetch":
			chain.filters = append(chain.filters, newFetchFilter())
		default:
			// Try to load plugin regarding the name
			plug := pr.LookupFilterPlugin(name)
			if plug == nil {
				return nil, fmt.Errorf("unsuported filter: %s", name)
			}
			chain.filters = append(chain.filters, plug.Filter)
		}
	}

	return chain, nil
}

// Apply applies filter chain on an article
func (c *Chain) Apply(article *model.Article) error {
	for idx, filter := range c.filters {
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
