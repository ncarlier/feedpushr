package store

import (
	"encoding/json"

	bolt "github.com/coreos/bbolt"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/common"
)

// OUTPUT_BUCKET bucket name
var OUTPUT_BUCKET = []byte("OUTPUT")

// GetOutput returns a stored Output.
func (store *BoltStore) GetOutput(ID int) (*app.Output, error) {
	var result app.Output
	err := store.get(OUTPUT_BUCKET, itob(ID), &result)
	if err != nil {
		if err == bolt.ErrInvalid {
			return nil, common.ErrOutputNotFound
		}
		return nil, err
	}
	return &result, nil
}

// DeleteOutput removes a output.
func (store *BoltStore) DeleteOutput(ID int) (*app.Output, error) {
	output, err := store.GetOutput(ID)
	if err != nil {
		return nil, err
	}

	err = store.delete(OUTPUT_BUCKET, itob(output.ID))
	if err != nil {
		return nil, err
	}
	return output, nil
}

// SaveOutput stores a output.
func (store *BoltStore) SaveOutput(output *app.Output) (*app.Output, error) {
	if output.ID == 0 {
		var err error
		id, err := store.nextSequence(OUTPUT_BUCKET)
		if err != nil {
			return nil, err
		}
		output.ID = int(id)
	}
	err := store.save(OUTPUT_BUCKET, itob(output.ID), &output)
	return output, err
}

// ListOutputs returns a paginated list of outputs.
func (store *BoltStore) ListOutputs(page, limit int) (*app.OutputCollection, error) {
	bufs, err := store.allAsRaw(OUTPUT_BUCKET, page, limit)
	if err != nil {
		return nil, err
	}

	outputs := app.OutputCollection{}
	for _, buf := range bufs {
		var output *app.Output
		if err := json.Unmarshal(buf, &output); err != nil {
			return nil, err
		}
		outputs = append(outputs, output)
	}
	return &outputs, nil
}

// ForEachOutput iterates over all outputs
func (store *BoltStore) ForEachOutput(cb func(*app.Output) error) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(OUTPUT_BUCKET).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var output *app.Output
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
