package store

import (
	"encoding/json"
	"fmt"
	"net/url"

	bolt "github.com/coreos/bbolt"
)

// BoltStore is a data store backed by BoltDB
type BoltStore struct {
	db *bolt.DB
}

// NewBoltStore creates a data store backed by BoltDB
func NewBoltStore(datasource *url.URL) (*BoltStore, error) {
	db, err := bolt.Open(datasource.Host+datasource.Path, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(FEED_BUCKET)
		if err != nil {
			return fmt.Errorf("could not create '%s' bucket: %v", FEED_BUCKET, err)
		}
		_, err = tx.CreateBucketIfNotExists(CACHE_BUCKET)
		if err != nil {
			return fmt.Errorf("could not create '%s' bucket: %v", CACHE_BUCKET, err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}

	return &BoltStore{
		db: db,
	}, nil
}

// Close the DB.
func (store *BoltStore) Close() error {
	return store.db.Close()
}

func (store *BoltStore) save(bucketName []byte, key string, dataStruct interface{}) error {
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
		return bucket.Put([]byte(key), encodedRecord)
	})
	return err
}

func (store *BoltStore) exists(bucketName []byte, key string) (bool, error) {
	result := false
	err := store.db.View(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket(bucketName)
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		// Retrieve the record
		v := b.Get([]byte(key))
		result = v != nil
		return nil
	})

	return result, err
}

func (store *BoltStore) get(bucketName []byte, key string, dataStruct interface{}) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket(bucketName)
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		// Retrieve the record
		v := b.Get([]byte(key))
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

func (store *BoltStore) delete(bucketName []byte, key string) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket(bucketName)
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		return b.Delete([]byte(key))
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
				break
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
