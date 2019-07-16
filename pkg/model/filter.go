package model

// FilterProps constain properties of a filter
type FilterProps map[string]interface{}

// Filter is the filter interface
type Filter interface {
	DoFilter(article *Article) error
	GetDef() FilterDef
}

// FilterDef contains filter definition
type FilterDef struct {
	ID int
	Spec
	Tags  []string
	Props FilterProps
}
