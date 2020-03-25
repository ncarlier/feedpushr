package store

import (
	"encoding/json"

	bolt "github.com/coreos/bbolt"
	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

// OUTPUT_BUCKET bucket name
var OUTPUT_BUCKET = []byte("OUTPUT")

// ClearOutputs clear all outputs
func (store *BoltStore) ClearOutputs() error {
	return store.clear(OUTPUT_BUCKET)
}

// GetOutput returns a stored Output.
func (store *BoltStore) GetOutput(ID string) (*model.OutputDef, error) {
	var result model.OutputDef
	err := store.get(OUTPUT_BUCKET, []byte(ID), &result)
	if err != nil {
		if err == bolt.ErrInvalid {
			return nil, common.ErrOutputNotFound
		}
		return nil, err
	}
	return &result, nil
}

// DeleteOutput removes a output.
func (store *BoltStore) DeleteOutput(ID string) (*model.OutputDef, error) {
	output, err := store.GetOutput(ID)
	if err != nil {
		return nil, err
	}

	err = store.delete(OUTPUT_BUCKET, []byte(output.ID))
	if err != nil {
		return nil, err
	}
	return output, nil
}

// SaveOutput stores a output.
func (store *BoltStore) SaveOutput(output model.OutputDef) (*model.OutputDef, error) {
	err := store.save(OUTPUT_BUCKET, []byte(output.ID), &output)
	return &output, err
}

// ListOutputs returns a paginated list of outputs.
func (store *BoltStore) ListOutputs(page, limit int) (*model.OutputDefCollection, error) {
	bufs, err := store.allAsRaw(OUTPUT_BUCKET, page, limit)
	if err != nil {
		return nil, err
	}

	outputs := model.OutputDefCollection{}
	for _, buf := range bufs {
		var output *model.OutputDef
		if err := json.Unmarshal(buf, &output); err != nil {
			return nil, err
		}
		outputs = append(outputs, output)
	}
	return &outputs, nil
}

// ForEachOutput iterates over all outputs
func (store *BoltStore) ForEachOutput(cb func(*model.OutputDef) error) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(OUTPUT_BUCKET).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var output *model.OutputDef
			if err := json.Unmarshal(v, &output); err != nil {
				// Unable to parse bucket payload
				output = nil
			}
			if err := cb(output); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
