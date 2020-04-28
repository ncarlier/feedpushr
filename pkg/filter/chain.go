package filter

import (
	"fmt"
	"sync"

	"github.com/ncarlier/feedpushr/v3/pkg/filter/plugins"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/plugin"
)

// Chain contains filter chain
type Chain struct {
	plugins map[string]model.FilterPlugin
	filters []model.Filter
	lock    sync.RWMutex
}

// NewChainFilter create new chain filter
func NewChainFilter(definitions model.FilterDefCollection) (*Chain, error) {
	chain := &Chain{
		plugins: plugins.GetBuiltinFilterPlugins(),
	}
	// Register external output plugins...
	err := plugin.GetRegistry().ForEachFilterPlugin(func(plug model.FilterPlugin) error {
		chain.plugins[plug.Spec().Name] = plug
		return nil
	})
	if err != nil {
		return nil, err
	}
	for _, def := range definitions {
		_, err := chain.Add(def)
		if err != nil {
			return nil, err
		}
	}
	return chain, nil
}

// GetAvailableFilters get all available filters
func (chain *Chain) GetAvailableFilters() []model.Spec {
	result := []model.Spec{}
	for _, plugin := range chain.plugins {
		result = append(result, plugin.Spec())
	}
	return result
}

// Apply applies filter chain on an article
func (chain *Chain) Apply(article *model.Article) error {
	for idx, filter := range chain.filters {
		if filter.GetDef().Enabled && filter.Match(article) {
			// TODO what to do with applied status?
			if _, err := filter.DoFilter(article); err != nil {
				return fmt.Errorf("error while applying filter #%d: %v", idx, err)
			}
		}
	}
	return nil
}
