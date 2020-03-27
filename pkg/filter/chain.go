package filter

import (
	"fmt"
	"sync"

	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
	"github.com/ncarlier/feedpushr/v2/pkg/plugin"
	"github.com/rs/zerolog/log"
)

// Chain contains filter chain
type Chain struct {
	filters []model.Filter
	lock    sync.RWMutex
}

// NewChainFilter create new chain filter
func NewChainFilter(definitions model.FilterDefCollection) (*Chain, error) {
	chain := &Chain{}
	for _, def := range definitions {
		_, err := chain.Add(def)
		if err != nil {
			return nil, err
		}
	}
	return chain, nil
}

func newFilter(def *model.FilterDef) (model.Filter, error) {
	var filter model.Filter
	var err error
	switch def.Name {
	case "title":
		filter, err = newTitleFilter(def)
	case "fetch":
		filter, err = newFetchFilter(def)
	case "minify":
		filter, err = newMinifyFilter(def)
	default:
		// Try to load plugin regarding the name
		plug := plugin.GetRegistry().LookupFilterPlugin(def.Name)
		if plug == nil {
			return nil, fmt.Errorf("unsupported filter: %s", def.Name)
		}
		filter, err = plug.Build(def)
	}
	if err != nil {
		return nil, fmt.Errorf("unable to create filter: %v", err)
	}
	return filter, nil
}

// GetAvailableFilters get all available filters
func GetAvailableFilters() []model.Spec {
	result := []model.Spec{
		titleSpec,
		fetchSpec,
		minifySpec,
	}
	plugin.GetRegistry().ForEachFilterPlugin(func(plug model.FilterPlugin) error {
		result = append(result, plug.Spec())
		return nil
	})
	return result
}

// Add a filter to the chain
func (chain *Chain) Add(filter *model.FilterDef) (model.Filter, error) {
	chain.lock.RLock()
	defer chain.lock.RUnlock()

	log.Debug().Str("id", filter.ID).Str("name", filter.Name).Msg("adding filter...")
	_filter, err := newFilter(filter)
	if err != nil {
		return nil, err
	}

	chain.filters = append(chain.filters, _filter)
	log.Info().Str("id", filter.ID).Str("name", filter.Name).Msg("filter added to the filter chain")
	return _filter, nil
}

// Update a filter of the chain
func (chain *Chain) Update(id string, update *model.FilterDef) (model.Filter, error) {
	chain.lock.RLock()
	defer chain.lock.RUnlock()

	for idx, f := range chain.filters {
		if f.GetDef().ID == id {
			update.ID = id
			update.Name = f.GetDef().Name
			log.Debug().Str("id", id).Str("name", update.Name).Msg("updating filter...")
			f, err := newFilter(update)
			if err != nil {
				return nil, err
			}
			chain.filters[idx] = f
			log.Info().Str("id", id).Str("name", update.Name).Msg("filter updated in the filter chain")
			return f, nil
		}
	}
	return nil, common.ErrFilterNotFound
}

// Remove a filter from the chain
func (chain *Chain) Remove(id string) error {
	chain.lock.RLock()
	defer chain.lock.RUnlock()

	for idx, f := range chain.filters {
		if f.GetDef().ID == id {
			log.Debug().Str("id", id).Str("name", f.GetDef().Name).Msg("removing filter...")
			chain.filters = append(chain.filters[:idx], chain.filters[idx+1:]...)
			log.Info().Str("id", id).Str("name", f.GetDef().Name).Msg("filter removed from filter chain")
			return nil
		}
	}
	return common.ErrFilterNotFound
}

// Apply applies filter chain on an article
func (chain *Chain) Apply(article *model.Article) error {
	for idx, filter := range chain.filters {
		if err := filter.DoFilter(article); err != nil {
			return fmt.Errorf("error while applying filter #%d: %v", idx, err)
		}
	}
	return nil
}

// Get returns a filter of the chain filter
func (chain *Chain) Get(id string) (model.Filter, error) {
	for _, f := range chain.filters {
		if f.GetDef().ID == id {
			return f, nil
		}
	}
	return nil, common.ErrFilterNotFound
}

// GetFilterDefs return definitions of the chain filter
func (chain *Chain) GetFilterDefs() model.FilterDefCollection {
	result := make(model.FilterDefCollection, len(chain.filters))
	for idx, filter := range chain.filters {
		def := filter.GetDef()
		result[idx] = &def
	}
	return result
}
