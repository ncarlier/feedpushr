package store

import (
	"github.com/ncarlier/feedpushr/autogen/app"
)

// OutputRepository interface to manage feeds
type OutputRepository interface {
	ListOutputs(page, limit int) (*app.OutputCollection, error)
	GetOutput(ID int) (*app.Output, error)
	DeleteOutput(ID int) (*app.Output, error)
	SaveOutput(filter *app.Output) (*app.Output, error)
	ForEachOutput(cb func(*app.Output) error) error
}
