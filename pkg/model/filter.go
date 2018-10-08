package model

// Filter is the filter interface
type Filter interface {
	DoFilter(article *Article) error
	GetSpec() FilterSpec
}

// FilterSpec contains filter specifications
type FilterSpec struct {
	Name  string
	Desc  string
	Tags  []string
	Props map[string]interface{}
}
