package model

// OutputProvider is the output provider interface
type OutputProvider interface {
	Send(article *Article) error
}
