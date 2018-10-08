package builder

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/ncarlier/feedpushr/autogen/app"
)

var tagRe = regexp.MustCompile("[^a-zA-Z0-9_]+")

// GetFeedID converts URL to feed ID (HASH)
func GetFeedID(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))
	return hex.EncodeToString(hasher.Sum(nil))
}

// NewFeed creates new Feed DTO
func NewFeed(url string, tags *string) (*app.Feed, error) {
	fp := gofeed.NewParser()
	fp.AtomTranslator = NewCustomAtomTranslator()
	fp.RSSTranslator = NewCustomRSSTranslator()

	rawFeed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	feed := &app.Feed{
		ID:     GetFeedID(url),
		XMLURL: url,
		Title:  rawFeed.Title,
		Mdate:  time.Now(),
		Cdate:  time.Now(),
		Tags:   GetFeedTags(tags),
	}

	if hub, ok := rawFeed.Custom["hub"]; ok {
		feed.HubURL = &hub
	}

	return feed, nil
}

// GetFeedTags extracts tags from a comma separated list of tags
func GetFeedTags(tags *string) []string {
	if tags == nil || strings.Trim(*tags, " ") == "" {
		return []string{}
	}

	result := strings.Split(*tags, ",")
	for i, v := range result {
		v = strings.TrimPrefix(v, "/")
		result[i] = tagRe.ReplaceAllString(v, "_")
	}
	return deduplicate(result)
}

func deduplicate(list []string) []string {
	keys := make(map[string]bool)
	result := []string{}
	for _, entry := range list {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}
