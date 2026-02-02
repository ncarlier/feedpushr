package store

import (
	"context"
	"encoding/json"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	bolt "go.etcd.io/bbolt"
)

// OutputBucketName bucket name
var OutputBucketName = []byte("OUTPUT")

// ClearOutputs clear all outputs
func (store *BoltStore) ClearOutputs(ctx context.Context) error {
	return store.clear(OutputBucketName)
}

// GetOutput returns a stored Output.
func (store *BoltStore) GetOutput(ctx context.Context, ID string) (*model.OutputDef, error) {
	var result model.OutputDef
	err := store.get(OutputBucketName, []byte(ID), &result)
	if err != nil {
		if err == bolt.ErrInvalid {
			return nil, common.ErrOutputNotFound
		}
		return nil, err
	}
	return &result, nil
}

// DeleteOutput removes a output.
func (store *BoltStore) DeleteOutput(ctx context.Context, ID string) (*model.OutputDef, error) {
	output, err := store.GetOutput(ctx, ID)
	if err != nil {
		return nil, err
	}

	err = store.delete(OutputBucketName, []byte(output.ID))
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (store *BoltStore) assertOutputQuota(ctx context.Context, output *model.OutputDef) error {
	if store.quota.MaxNbOutputs > 0 {
		if exists, err := store.exists(OutputBucketName, []byte(output.ID)); err != nil {
			return err
		} else if !exists {
			total, err := store.count(OutputBucketName)
			if err != nil {
				return err
			}
			if total >= store.quota.MaxNbOutputs {
				return common.ErrOutputQuotaExceeded
			}
		}
	}
	return nil
}

// SaveOutput stores a output.
func (store *BoltStore) SaveOutput(ctx context.Context, output model.OutputDef) (*model.OutputDef, error) {
	if err := store.assertOutputQuota(ctx, &output); err != nil {
		return nil, err
	}
	err := store.save(OutputBucketName, []byte(output.ID), &output)
	return &output, err
}

// ListOutputs returns a paginated list of outputs.
func (store *BoltStore) ListOutputs(ctx context.Context, page, limit int) (*model.OutputDefCollection, error) {
	bufs, err := store.allAsRaw(OutputBucketName, page, limit)
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
func (store *BoltStore) ForEachOutput(ctx context.Context, cb func(*model.OutputDef) error) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(OutputBucketName).Cursor()
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
