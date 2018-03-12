package model

import (
	"time"
)

// CacheItem stored into the cache.
type CacheItem struct {
	Value string    `json:"value"`
	Date  time.Time `json:"date"`
}
