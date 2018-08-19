package model

// OutputProvider is the output provider interface
type OutputProvider interface {
	Send(article *Article) error
	GetSpec() OutputSpec
}

// OutputSpec contains output specifications
type OutputSpec struct {
	Name  string
	Desc  string
	Props map[string]interface{}
}
