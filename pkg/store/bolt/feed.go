package store

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	bolt "github.com/coreos/bbolt"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/common"
)

// FEED_BUCKET bucket name
var FEED_BUCKET = []byte("FEED")

// ExistsFeed returns true if a feed exists for this url.
func (store *BoltStore) ExistsFeed(url string) bool {
	hasher := md5.New()
	hasher.Write([]byte(url))
	id := hex.EncodeToString(hasher.Sum(nil))

	exists, err := store.exists(FEED_BUCKET, []byte(id))
	if err != nil {
		return false
	}
	return exists
}

// GetFeed returns a stored Feed.
func (store *BoltStore) GetFeed(id string) (*app.Feed, error) {
	var result app.Feed
	err := store.get(FEED_BUCKET, []byte(id), &result)
	if err != nil {
		if err == bolt.ErrInvalid {
			return nil, common.ErrFeedNotFound
		}
		return nil, err
	}
	return &result, nil
}

// DeleteFeed removes a feed.
func (store *BoltStore) DeleteFeed(id string) (*app.Feed, error) {
	feed, err := store.GetFeed(id)
	if err != nil {
		return nil, err
	}

	err = store.delete(FEED_BUCKET, []byte(feed.ID))
	if err != nil {
		return nil, err
	}
	return feed, nil
}

// SaveFeed stores a feed.
func (store *BoltStore) SaveFeed(feed *app.Feed) error {
	return store.save(FEED_BUCKET, []byte(feed.ID), &feed)
}

// ListFeeds returns a paginated list of feeds.
func (store *BoltStore) ListFeeds(page, limit int) (*app.FeedCollection, error) {
	bufs, err := store.allAsRaw(FEED_BUCKET, page, limit)
	if err != nil {
		return nil, err
	}

	feeds := app.FeedCollection{}
	for _, buf := range bufs {
		var feed *app.Feed
		if err := json.Unmarshal(buf, &feed); err != nil {
			return nil, err
		}
		feeds = append(feeds, feed)
	}
	return &feeds, nil
}

// CountFeeds returns total numer of feeds.
func (store *BoltStore) CountFeeds() (int, error) {
	return store.count(FEED_BUCKET)
}

// ForEachFeed iterates over all feeds
func (store *BoltStore) ForEachFeed(cb func(*app.Feed) error) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(FEED_BUCKET).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var feed *app.Feed
			if err := json.Unmarshal(v, &feed); err != nil {
				// Unable to parse bucket payload
				feed = nil
			}
			if err := cb(feed); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
