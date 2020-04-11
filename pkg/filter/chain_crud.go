package filter

import (
	"fmt"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/rs/zerolog/log"
)

// Add a filter to the chain
func (chain *Chain) Add(def *model.FilterDef) (model.Filter, error) {
	chain.lock.RLock()
	defer chain.lock.RUnlock()

	log.Debug().Str("id", def.ID).Str("name", def.Name).Msg("adding filter...")
	plug, ok := chain.plugins[def.Name]
	if !ok {
		return nil, fmt.Errorf("unsupported filter: %s", def.Name)
	}

	filter, err := plug.Build(def)
	if err != nil {
		return nil, err
	}

	chain.filters = append(chain.filters, filter)
	log.Info().Str("id", def.ID).Str("name", def.Name).Msg("filter added to the filter chain")
	return filter, nil
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
			plug, ok := chain.plugins[update.Name]
			if !ok {
				return nil, fmt.Errorf("unsupported filter: %s", update.Name)
			}
			f, err := plug.Build(update)
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
