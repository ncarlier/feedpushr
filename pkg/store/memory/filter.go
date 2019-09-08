package store

import (
	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/model"
)

func (store *InMemoryStore) nextFilterSequence() int {
	max := 0
	for _, filter := range store.filters {
		if filter.ID > max {
			max = filter.ID
		}
	}
	return max + 1
}

// ClearFilters clear all filters
func (store *InMemoryStore) ClearFilters() error {
	store.filtersLock.RLock()
	defer store.filtersLock.RUnlock()
	store.filters = make(map[int]model.FilterDef)
	return nil
}

// GetFilter returns a stored Filter.
func (store *InMemoryStore) GetFilter(ID int) (*model.FilterDef, error) {
	filter, exists := store.filters[ID]
	if !exists {
		return nil, common.ErrFilterNotFound
	}
	return &filter, nil
}

// DeleteFilter removes a filter.
func (store *InMemoryStore) DeleteFilter(ID int) (*model.FilterDef, error) {
	store.filtersLock.RLock()
	defer store.filtersLock.RUnlock()
	filter, err := store.GetFilter(ID)
	if err != nil {
		return nil, err
	}
	delete(store.filters, ID)
	return filter, nil
}

// SaveFilter stores a filter.
func (store *InMemoryStore) SaveFilter(filter model.FilterDef) (*model.FilterDef, error) {
	store.filtersLock.RLock()
	defer store.filtersLock.RUnlock()
	if filter.ID == 0 {
		filter.ID = store.nextFilterSequence()
	}
	store.filters[filter.ID] = filter
	return &filter, nil
}

// ListFilters returns a paginated list of filters.
func (store *InMemoryStore) ListFilters(page, limit int) (*model.FilterDefCollection, error) {
	filters := model.FilterDefCollection{}
	startOffset := (page - 1) * limit
	offset := 0
	for _, filter := range store.filters {
		switch {
		case offset < startOffset:
			// Skip entries before the start offset
			offset++
			continue
		case offset >= startOffset+limit:
			// End of the window
			break
		default:
			// Add value to entries
			offset++
			filters = append(filters, &filter)
		}
	}
	return &filters, nil
}

// ForEachFilter iterates over all filters
func (store *InMemoryStore) ForEachFilter(cb func(*model.FilterDef) error) error {
	store.filtersLock.RLock()
	defer store.filtersLock.RUnlock()
	for _, filter := range store.filters {
		if err := cb(&filter); err != nil {
			return err
		}
	}
	return nil
}
