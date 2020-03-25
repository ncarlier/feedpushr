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
func NewBoltStore(datasource *url.URL) (*BoltStore, error) {
	db, err := bolt.Open(datasource.Host+datasource.Path, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		return createBucketsIfNotExists(
			tx,
			FEED_BUCKET,
			OUTPUT_BUCKET,
			CACHE_BUCKET,
		)
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

func (store *BoltStore) nextSequence(bucketName []byte) (uint64, error) {
	var result uint64
	err := store.db.Update(func(tx *bolt.Tx) error {
		bucket, e := tx.CreateBucketIfNotExists(bucketName)
		if e != nil {
			return e
		}
		result, e = bucket.NextSequence()
		return e
	})
	return result, err

}
