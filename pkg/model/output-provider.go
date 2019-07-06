package model

import (
	"crypto/md5"
	"encoding/hex"
)

// OutputProps contains properties of an output
type OutputProps map[string]interface{}

// OutputProvider is the output provider interface
type OutputProvider interface {
	Send(article *Article) error
	GetSpec() OutputSpec
}

// OutputSpec contains output specifications
type OutputSpec struct {
	Name  string
	Desc  string
	Tags  []string
	Props map[string]interface{}
}

// Hash computes spec hash
func (spec OutputSpec) Hash() string {
	// TODO add props to the key computation
	key := spec.Name
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
