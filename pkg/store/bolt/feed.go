package store

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	bolt "github.com/coreos/bbolt"
	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// FeedBucketName bucket name
var FeedBucketName = []byte("FEED")

// ExistsFeed returns true if a feed exists for this url.
func (store *BoltStore) ExistsFeed(url string) bool {
	hasher := md5.New()
	hasher.Write([]byte(url))
	id := hex.EncodeToString(hasher.Sum(nil))

	exists, err := store.exists(FeedBucketName, []byte(id))
	if err != nil {
		return false
	}
	return exists
}

// GetFeed returns a stored Feed.
func (store *BoltStore) GetFeed(id string) (*model.FeedDef, error) {
	var result model.FeedDef
	err := store.get(FeedBucketName, []byte(id), &result)
	if err != nil {
		if err == bolt.ErrInvalid {
			return nil, common.ErrFeedNotFound
		}
		return nil, err
	}
	return &result, nil
}

// DeleteFeed removes a feed.
func (store *BoltStore) DeleteFeed(id string) (*model.FeedDef, error) {
	feed, err := store.GetFeed(id)
	if err != nil {
		return nil, err
	}

	err = store.delete(FeedBucketName, []byte(feed.ID))
	if err != nil {
		return nil, err
	}
	err = store.index.Delete(feed.ID)
	if err != nil {
		return nil, err
	}
	return feed, nil
}

// SaveFeed stores a feed.
func (store *BoltStore) SaveFeed(feed *model.FeedDef) error {
	err := store.save(FeedBucketName, []byte(feed.ID), &feed)
	if err != nil {
		return err
	}
	return store.index.Index(feed.ID, feed)
}

// ListFeeds returns a paginated list of feeds.
func (store *BoltStore) ListFeeds(page, size int) (*model.FeedDefPage, error) {
	bufs, err := store.allAsRaw(FeedBucketName, page, size)
	if err != nil {
		return nil, err
	}

	total, err := store.CountFeeds()
	if err != nil {
		return nil, err
	}

	result := model.FeedDefPage{
		Page:  page,
		Size:  size,
		Total: total,
	}
	for _, buf := range bufs {
		var feed model.FeedDef
		if err := json.Unmarshal(buf, &feed); err != nil {
			return nil, err
		}
		result.Feeds = append(result.Feeds, feed)
	}
	return &result, nil
}

// CountFeeds returns total numer of feeds.
func (store *BoltStore) CountFeeds() (int, error) {
	return store.count(FeedBucketName)
}

// ForEachFeed iterates over all feeds
func (store *BoltStore) ForEachFeed(cb func(*model.FeedDef) error) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(FeedBucketName).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var feed model.FeedDef
			if err := json.Unmarshal(v, &feed); err == nil {
				// Unable to parse bucket payload
				if err := cb(&feed); err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}
