package store

import (
	"github.com/ncarlier/feedpushr/pkg/model"
)

// FilterRepository interface to manage feeds
type FilterRepository interface {
	ListFilters(page, limit int) (*model.FilterDefCollection, error)
	GetFilter(ID int) (*model.FilterDef, error)
	DeleteFilter(ID int) (*model.FilterDef, error)
	SaveFilter(filter model.FilterDef) (*model.FilterDef, error)
	ForEachFilter(cb func(*model.FilterDef) error) error
	ClearFilters() error
}
