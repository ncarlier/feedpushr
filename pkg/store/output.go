package store

import (
	"context"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// OutputRepository interface to manage feeds
type OutputRepository interface {
	ListOutputs(ctx context.Context, page, limit int) (*model.OutputDefCollection, error)
	GetOutput(ctx context.Context, ID string) (*model.OutputDef, error)
	DeleteOutput(ctx context.Context, ID string) (*model.OutputDef, error)
	SaveOutput(ctx context.Context, output model.OutputDef) (*model.OutputDef, error)
	ForEachOutput(ctx context.Context, cb func(*model.OutputDef) error) error
	ClearOutputs(ctx context.Context) error
}
