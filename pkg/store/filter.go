package store

import (
	"github.com/ncarlier/feedpushr/autogen/app"
)

// FilterRepository interface to manage feeds
type FilterRepository interface {
	ListFilters(page, limit int) (*app.FilterCollection, error)
	GetFilter(ID int) (*app.Filter, error)
	DeleteFilter(ID int) (*app.Filter, error)
	SaveFilter(filter *app.Filter) (*app.Filter, error)
	ForEachFilter(cb func(*app.Filter) error) error
}
