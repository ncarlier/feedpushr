package model

import "time"

// Article model structure.
type Article struct {
	Title           string     `json:"title,omitempty"`
	Description     string     `json:"description,omitempty"`
	Content         string     `json:"content,omitempty"`
	Link            string     `json:"link,omitempty"`
	Updated         string     `json:"updated,omitempty"`
	UpdatedParsed   *time.Time `json:"updatedParsed,omitempty"`
	Published       string     `json:"published,omitempty"`
	PublishedParsed *time.Time `json:"publishedParsed,omitempty"`
	GUID            string     `json:"guid,omitempty"`
}
