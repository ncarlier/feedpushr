package store

import (
	"encoding/binary"
	"encoding/json"

	bolt "github.com/coreos/bbolt"
	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/model"
)

// FILTER_BUCKET bucket name
var FILTER_BUCKET = []byte("FILTER")

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// GetFilter returns a stored Filter.
func (store *BoltStore) GetFilter(ID int) (*model.FilterDef, error) {
	var result model.FilterDef
	err := store.get(FILTER_BUCKET, itob(ID), &result)
	if err != nil {
		if err == bolt.ErrInvalid {
			return nil, common.ErrFilterNotFound
		}
		return nil, err
	}
	return &result, nil
}

// DeleteFilter removes a filter.
func (store *BoltStore) DeleteFilter(ID int) (*model.FilterDef, error) {
	filter, err := store.GetFilter(ID)
	if err != nil {
		return nil, err
	}

	err = store.delete(FILTER_BUCKET, itob(filter.ID))
	if err != nil {
		return nil, err
	}
	return filter, nil
}

// SaveFilter stores a filter.
func (store *BoltStore) SaveFilter(filter model.FilterDef) (*model.FilterDef, error) {
	if filter.ID == 0 {
		var err error
		id, err := store.nextSequence(FILTER_BUCKET)
		if err != nil {
			return nil, err
		}
		filter.ID = int(id)
	}
	err := store.save(FILTER_BUCKET, itob(filter.ID), &filter)
	return &filter, err
}

// ListFilters returns a paginated list of filters.
func (store *BoltStore) ListFilters(page, limit int) (*model.FilterDefCollection, error) {
	bufs, err := store.allAsRaw(FILTER_BUCKET, page, limit)
	if err != nil {
		return nil, err
	}

	filters := model.FilterDefCollection{}
	for _, buf := range bufs {
		var filter *model.FilterDef
		if err := json.Unmarshal(buf, &filter); err != nil {
			return nil, err
		}
		filters = append(filters, filter)
	}
	return &filters, nil
}

// ForEachFilter iterates over all filters
func (store *BoltStore) ForEachFilter(cb func(*model.FilterDef) error) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(FILTER_BUCKET).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var filter *model.FilterDef
			if err := json.Unmarshal(v, &filter); err != nil {
				// Unable to parse bucket payload
				filter = nil
			}
			if err := cb(filter); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
