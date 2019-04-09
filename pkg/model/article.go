package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// Article model structure.
type Article struct {
	Title           string                 `json:"title,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Content         string                 `json:"content,omitempty"`
	Link            string                 `json:"link,omitempty"`
	Updated         string                 `json:"updated,omitempty"`
	UpdatedParsed   *time.Time             `json:"updatedParsed,omitempty"`
	Published       string                 `json:"published,omitempty"`
	PublishedParsed *time.Time             `json:"publishedParsed,omitempty"`
	GUID            string                 `json:"guid,omitempty"`
	Meta            map[string]interface{} `json:"meta,omitempty"`
	Tags            []string               `json:"tags,omitempty"`
}

func (a *Article) String() string {
	result, _ := json.Marshal(a)
	return string(result)
}

// RefDate get article reference date (published or updated date)
func (a *Article) RefDate() *time.Time {
	var date *time.Time
	if a.PublishedParsed != nil {
		date = a.PublishedParsed
	}
	if a.UpdatedParsed != nil {
		date = a.UpdatedParsed
	}
	return date
}

// IsValid test if the article can be pushed
func (a *Article) IsValid(maxAge time.Time) error {
	date := a.RefDate()
	if date == nil {
		return fmt.Errorf("missing article date")
	}
	if date.Before(maxAge) {
		return fmt.Errorf("article too old")
	}
	return nil
}

// Hash computes article hash
func (a *Article) Hash() string {
	key := a.GUID
	if key == "" {
		key = a.Link
	}
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Match test if articles tags matches other tags
func (a *Article) Match(tags []string) bool {
	// If no tags are provided then the article match
	if len(tags) == 0 {
		return true
	}
	tagSet := make(map[string]struct{}, len(a.Tags))
	for _, tag := range a.Tags {
		tagSet[tag] = struct{}{}
	}

	for _, tag := range tags {
		if _, ok := tagSet[tag]; !ok {
			return false
		}
	}
	return true
}
