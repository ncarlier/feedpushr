package store

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/blevesearch/bleve"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

// BoltStore is a data store backed by BoltDB
type BoltStore struct {
	db    *bolt.DB
	index bleve.Index
	quota model.Quota
}

func createBucketsIfNotExists(tx *bolt.Tx, buckets ...[]byte) error {
	for _, bucket := range buckets {
		_, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return fmt.Errorf("could not create '%s' bucket: %v", bucket, err)
		}
	}
	return nil
}

// NewBoltStore creates a data store backed by BoltDB
func NewBoltStore(datasource *url.URL, quota model.Quota) (*BoltStore, error) {
	dbPath := datasource.Host + datasource.Path
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to open DB, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		return createBucketsIfNotExists(
			tx,
			FeedBucketName,
			OutputBucketName,
			CacheBucketName,
		)
	})
	if err != nil {
		return nil, fmt.Errorf("unable to set up buckets, %v", err)
	}

	// Init search index
	index, err := openSearchIndex(db.Path())
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("unable to open search index, %v", err)
	}

	return &BoltStore{
		db:    db,
		index: index,
		quota: quota,
	}, nil
}

// Close the DB.
func (store *BoltStore) Close() error {
	var err error
	if err = store.index.Close(); err != nil {
		if err2 := store.db.Close(); err2 != nil {
			err = errors.Wrap(err, err2.Error())
		}
	} else {
		err = store.db.Close()
	}
	return err
}

func (store *BoltStore) clear(bucketName []byte) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		// Remove the bucket
		if e := tx.DeleteBucket(bucketName); e != nil {
			return e
		}
		// Create the bucket
		_, e := tx.CreateBucketIfNotExists(bucketName)
		return e
	})
	return err
}

func (store *BoltStore) count(bucketName []byte) (nb int, err error) {
	nb = 0
	err = store.db.View(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket(bucketName)
		if b == nil {
			return bolt.ErrBucketNotFound
		}
		stats := b.Stats()
		nb = stats.KeyN
		return nil
	})
	return
}

func (store *BoltStore) save(bucketName, key []byte, dataStruct interface{}) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		// Create the bucket
		bucket, e := tx.CreateBucketIfNotExists(bucketName)
		if e != nil {
			return e
		}

		// Encode the record
		encodedRecord, e := json.Marshal(dataStruct)
		if e != nil {
			return e
		}

		// Store the record
		return bucket.Put(key, encodedRecord)
	})
	return err
}

func (store *BoltStore) exists(bucketName, key []byte) (bool, error) {
	result := false
	err := store.db.View(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket(bucketName)
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		// Retrieve the record
		v := b.Get(key)
		result = v != nil
		return nil
	})

	return result, err
}

func (store *BoltStore) get(bucketName, key []byte, dataStruct interface{}) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket(bucketName)
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		// Retrieve the record
		v := b.Get(key)
		if len(v) < 1 {
			return bolt.ErrInvalid
		}

		// Decode the record
		e := json.Unmarshal(v, &dataStruct)
		if e != nil {
			return e
		}

		return nil
	})

	return err
}

func (store *BoltStore) delete(bucketName, key []byte) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket(bucketName)
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		return b.Delete(key)
	})
	return err
}

func (store *BoltStore) allAsRaw(bucket []byte, page, size int) ([][]byte, error) {
	entries := [][]byte{}
	startOffset := (page - 1) * size
	err := store.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucket).Cursor()
		offset := 0
		for k, v := c.First(); k != nil; k, v = c.Next() {
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
				entries = append(entries, v)
			}
		}
		return nil
	})
	return entries, err
}
