package store

import (
	"github.com/ncarlier/feedpushr/pkg/model"
)

// OutputRepository interface to manage feeds
type OutputRepository interface {
	ListOutputs(page, limit int) (*model.OutputDefCollection, error)
	GetOutput(ID int) (*model.OutputDef, error)
	DeleteOutput(ID int) (*model.OutputDef, error)
	SaveOutput(output model.OutputDef) (*model.OutputDef, error)
	ForEachOutput(cb func(*model.OutputDef) error) error
}
