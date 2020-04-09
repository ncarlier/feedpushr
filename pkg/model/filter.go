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
	// DoFilter apply filter on the article.
	// Returns true if the filter was applied
	DoFilter(article *Article) (bool, error)
	// Match test if article match with filter condition
	Match(article *Article) bool
	// GetDef returns filter definition
	GetDef() FilterDef
}

// FilterDefCollection is an array of filter definition
type FilterDefCollection []*FilterDef

// FilterDef contains filter definition
type FilterDef struct {
	ID    string `json:"id"`
	Alias string `json:"alias"`
	Spec
	Condition string      `json:"condition"`
	Props     FilterProps `json:"props:omitempty"`
	Enabled   bool        `json:"enabled"`
	NbSuccess uint64      `json:"nbSuccess"`
	NbError   uint64      `json:"nbError"`
}
