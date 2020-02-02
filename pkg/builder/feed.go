package builder

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"mime"
	"net/http"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/html"
	"github.com/ncarlier/feedpushr/v2/pkg/strcase"
)

// GetFeedID converts URL to feed ID (HASH)
func GetFeedID(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))
	return hex.EncodeToString(hasher.Sum(nil))
}

// NewFeed creates new Feed DTO
func NewFeed(url string, tags *string) (*app.Feed, error) {
	// Set timeout context
	ctx, cancel := context.WithCancel(context.TODO())
	timeout := time.AfterFunc(common.DefaultTimeout, func() {
		cancel()
	})

	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", common.UserAgent)

	// Do HTTP call
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	timeout.Stop()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("http error: %s", res.Status)
	}

	// Get content-type
	contentTypeHeader := res.Header.Get("Content-type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		return nil, err
	}

	if contentType == "text/html" {
		urls, err := html.ExtractFeedLinks(res.Body)
		if err != nil {
			return nil, err
		}
		if len(urls) == 0 {
			return nil, fmt.Errorf("no feed URL found on this page: %s", url)
		}
		return NewFeed(urls[0], tags)
	}

	if !common.ValidFeedContentType.MatchString(contentType) {
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}

	// Parse feed
	fp := gofeed.NewParser()
	fp.AtomTranslator = NewCustomAtomTranslator()
	fp.RSSTranslator = NewCustomRSSTranslator()

	rawFeed, err := fp.Parse(res.Body)
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
