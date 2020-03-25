package model

import "fmt"

// FilterProps constain properties of a filter
type FilterProps map[string]interface{}

// Get property string value
func (p FilterProps) Get(key string) string {
	if val, ok := p[key]; ok {
		return fmt.Sprintf("%v", val)
	}
	return ""
}

// Filter is the filter interface
type Filter interface {
	DoFilter(article *Article) error
	GetDef() FilterDef
}

// FilterDefCollection is an array of filter definition
type FilterDefCollection []*FilterDef

// FilterDef contains filter definition
type FilterDef struct {
	Alias string `json:"alias"`
	Spec
	Condition string      `json:"condition"`
	Props     FilterProps `json:"props:omitempty"`
	Enabled   bool        `json:"enabled"`
	NbSuccess uint64      `json:"nbSuccess"`
	NbError   uint64      `json:"nbError"`
}
