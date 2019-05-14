package builder

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"

	"github.com/ncarlier/feedpushr/pkg/strcase"

	"github.com/mmcdole/gofeed"
	"github.com/ncarlier/feedpushr/autogen/app"
)

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

// JoinTags join tags in a comma separated string
func JoinTags(tags ...string) string {
	result := ""
	for _, tag := range tags {
		if result != "" {
			if tag != "" {
				result += "," + tag
			}
		} else {
			result = tag
		}
	}
	return result

}

// GetFeedTags extracts tags from a comma separated list of tags
func GetFeedTags(tags *string) []string {
	if tags == nil || strings.Trim(*tags, " ") == "" {
		return []string{}
	}

	result := strings.Split(*tags, ",")
	for i, v := range result {
		v = strings.TrimPrefix(v, "/")
		result[i] = strcase.ToSnake(v)
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
