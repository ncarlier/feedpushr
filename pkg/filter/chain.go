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
func NewChainFilter() *Chain {
	return &Chain{}
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

	log.Debug().Str("name", filter.Name).Msg("adding filter...")
	nextID := 0
	for _, _filter := range chain.filters {
		if nextID < _filter.GetDef().ID {
			nextID = _filter.GetDef().ID
		}
	}
	filter.ID = nextID + 1
	_filter, err := newFilter(filter)
	if err != nil {
		return nil, err
	}

	chain.filters = append(chain.filters, _filter)
	log.Info().Int("id", filter.ID).Str("name", filter.Name).Msg("filter added to the filter chain")
	return _filter, nil
}

// Update a filter of the chain
func (chain *Chain) Update(update *model.FilterDef) (model.Filter, error) {
	chain.lock.RLock()
	defer chain.lock.RUnlock()

	for idx, filter := range chain.filters {
		if update.ID == filter.GetDef().ID {
			log.Debug().Int("id", update.ID).Msg("updating filter...")
			update.Name = filter.GetDef().Name
			f, err := newFilter(update)
			if err != nil {
				return nil, err
			}
			chain.filters[idx] = f
			log.Info().Int("id", update.ID).Msg("filter updated")
			return f, nil
		}
	}
	return nil, common.ErrFilterNotFound
}

// Remove a filter from the chain
func (chain *Chain) Remove(filter *model.FilterDef) error {
	chain.lock.RLock()
	defer chain.lock.RUnlock()

	for idx, _filter := range chain.filters {
		if filter.ID == _filter.GetDef().ID {
			log.Debug().Int("id", filter.ID).Msg("removing filter...")
			chain.filters = append(chain.filters[:idx], chain.filters[idx+1:]...)
			log.Info().Int("id", filter.ID).Msg("filter removed from the filter chain")
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

// Get a filter from the chain
func (chain *Chain) Get(id int) (model.Filter, error) {
	for _, _filter := range chain.filters {
		if id == _filter.GetDef().ID {
			return _filter, nil
		}
	}
	return nil, common.ErrFilterNotFound
}

// GetFilterDefs return definitions of the chain filter
func (chain *Chain) GetFilterDefs() []model.FilterDef {
	result := make([]model.FilterDef, len(chain.filters))
	for idx, filter := range chain.filters {
		result[idx] = filter.GetDef()
	}
	return result
}
