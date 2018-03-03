package cache

import (
	"time"
)

// Item stored into the cache.
type Item struct {
	Value string    `json:"value"`
	Date  time.Time `json:"date"`
}
