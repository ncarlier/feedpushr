package model

// FilterProps constain properties of a filter
type FilterProps map[string]interface{}

// Filter is the filter interface
type Filter interface {
	DoFilter(article *Article) error
	GetSpec() FilterSpec
}

// FilterSpec contains filter specifications
type FilterSpec struct {
	ID    int
	Name  string
	Desc  string
	Tags  []string
	Props FilterProps
}
