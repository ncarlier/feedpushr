package model

import (
	"encoding/json"
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
