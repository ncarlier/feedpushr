package model

// Filter is the filter interface
type Filter interface {
	DoFilter(article *Article) error
}
