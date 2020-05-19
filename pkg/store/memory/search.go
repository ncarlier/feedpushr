package store

import "github.com/ncarlier/feedpushr/v3/pkg/model"

// BuildInitialIndex create initial index (only if empty)
func (store *InMemoryStore) BuildInitialIndex() error {
	// NOT IMPLEMENTED
	return nil
}

// SearchFeeds search feeds using search index
func (store *InMemoryStore) SearchFeeds(query string, page, size int) (*model.FeedDefPage, error) {
	result := model.FeedDefPage{
		Page: page,
		Size: size,
	}
	return &result, nil
}
