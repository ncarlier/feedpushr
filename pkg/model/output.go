package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// OutputProps contains properties of an output
type OutputProps map[string]interface{}

// Get property string value
func (p OutputProps) Get(key string) string {
	if val, ok := p[key]; ok {
		return fmt.Sprintf("%v", val)
	}
	return ""
}

// Output is the output interface
type Output interface {
	// Send article to the output.
	// Returns true if the article was sent
	Send(article *Article) (bool, error)
	// GetDef returns output definition
	GetDef() OutputDef
}

// OutputDefCollection is an array of output definition
type OutputDefCollection []*OutputDef

// OutputDef contains output definition
type OutputDef struct {
	ID    string `json:"id"`
	Alias string `json:"alias"`
	Spec
	Condition string              `json:"condition"`
	Props     OutputProps         `json:"props:omitempty"`
	Filters   FilterDefCollection `json:"filters"`
	Enabled   bool                `json:"enabled"`
	NbSuccess uint32              `json:"nbSuccess"`
	NbError   uint32              `json:"nbError"`
}

// Hash computes spec hash
func (def OutputDef) Hash() string {
	hasher := md5.New()
	hasher.Write([]byte(def.ID))
	return hex.EncodeToString(hasher.Sum(nil))
}
