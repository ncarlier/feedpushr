package store

import (
	"context"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// ClearOutputs clear all outputs
func (store *InMemoryStore) ClearOutputs(ctx context.Context) error {
	store.outputsLock.RLock()
	defer store.outputsLock.RUnlock()
	store.outputs = make(map[string]model.OutputDef)
	return nil
}

// GetOutput returns a stored Output.
func (store *InMemoryStore) GetOutput(ctx context.Context, ID string) (*model.OutputDef, error) {
	output, exists := store.outputs[ID]
	if !exists {
		return nil, common.ErrOutputNotFound
	}
	return &output, nil
}

// DeleteOutput removes a output.
func (store *InMemoryStore) DeleteOutput(ctx context.Context, ID string) (*model.OutputDef, error) {
	store.outputsLock.RLock()
	defer store.outputsLock.RUnlock()
	output, err := store.GetOutput(ctx, ID)
	if err != nil {
		return nil, err
	}
	delete(store.outputs, ID)
	return output, nil
}

// SaveOutput stores a output.
func (store *InMemoryStore) SaveOutput(ctx context.Context, output model.OutputDef) (*model.OutputDef, error) {
	store.outputsLock.RLock()
	defer store.outputsLock.RUnlock()
	store.outputs[output.ID] = output
	return &output, nil
}

// ListOutputs returns a paginated list of outputs.
func (store *InMemoryStore) ListOutputs(ctx context.Context, page, limit int) (*model.OutputDefCollection, error) {
	outputs := model.OutputDefCollection{}
	startOffset := (page - 1) * limit
	offset := 0
	for _, output := range store.outputs {
		switch {
		case offset < startOffset:
			// Skip entries before the start offset
			offset++
			continue
		case offset >= startOffset+limit:
			// End of the window
		default:
			// Add value to entries
			offset++
			outputs = append(outputs, &output)
		}
	}
	return &outputs, nil
}

// ForEachOutput iterates over all outputs
func (store *InMemoryStore) ForEachOutput(ctx context.Context, cb func(*model.OutputDef) error) error {
	store.outputsLock.RLock()
	defer store.outputsLock.RUnlock()
	for _, output := range store.outputs {
		if err := cb(&output); err != nil {
			return err
		}
	}
	return nil
}
