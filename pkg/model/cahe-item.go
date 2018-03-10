package model

import (
	"time"
)

// Item stored into the cache.
type CacheItem struct {
	Value string    `json:"value"`
	Date  time.Time `json:"date"`
}
