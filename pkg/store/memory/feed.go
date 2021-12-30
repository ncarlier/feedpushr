package store

import (
	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/helper"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// ExistsFeed returns true if a feed exists for this url.
func (store *InMemoryStore) ExistsFeed(url string) bool {
	id := helper.Hash(url)
	_, exists := store.feeds[id]
	return exists
}

// GetFeed returns a stored Feed.
func (store *InMemoryStore) GetFeed(id string) (*model.FeedDef, error) {
	feed, exists := store.feeds[id]
	if !exists {
		return nil, common.ErrFeedNotFound
	}
	return &feed, nil
}

// DeleteFeed removes a feed.
func (store *InMemoryStore) DeleteFeed(id string) (*model.FeedDef, error) {
	store.feedsLock.RLock()
	defer store.feedsLock.RUnlock()
	feed, err := store.GetFeed(id)
	if err != nil {
		return nil, err
	}
	delete(store.feeds, id)
	return feed, nil
}

// SaveFeed stores a feed.
func (store *InMemoryStore) SaveFeed(feed *model.FeedDef) error {
	store.feedsLock.RLock()
	defer store.feedsLock.RUnlock()
	store.feeds[feed.ID] = *feed
	return nil
}

// ListFeeds returns a paginated list of feeds.
func (store *InMemoryStore) ListFeeds(page, size int) (*model.FeedDefPage, error) {
	if page <= 0 {
		page = 1
	}
	result := model.FeedDefPage{
		Page:  page,
		Size:  size,
		Total: len(store.feeds),
	}
	startOffset := (page - 1) * size
	offset := 0
	for _, feed := range store.feeds {
		switch {
		case offset < startOffset:
			// Skip entries before the start offset
			offset++
			continue
		case offset >= startOffset+size:
			// End of the window
		default:
			// Add value to entries
			offset++
			result.Feeds = append(result.Feeds, feed)
		}
	}
	return &result, nil
}

// CountFeeds returns total numer of feeds.
func (store *InMemoryStore) CountFeeds() (int, error) {
	return len(store.feeds), nil
}

// ForEachFeed iterates over all feeds
func (store *InMemoryStore) ForEachFeed(cb func(*model.FeedDef) error) error {
	store.feedsLock.RLock()
	defer store.feedsLock.RUnlock()
	for _, feed := range store.feeds {
		if err := cb(&feed); err != nil {
			return err
		}
	}
	return nil
}
