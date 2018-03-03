package builder

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/ncarlier/feedpushr/autogen/app"
)

// NewFeed creates new Feed DTO
func NewFeed(url string) (*app.Feed, error) {
	fp := gofeed.NewParser()
	fp.AtomTranslator = NewCustomAtomTranslator()
	fp.RSSTranslator = NewCustomRSSTranslator()

	rawFeed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	feed := &app.Feed{
		XMLURL: url,
		Title:  rawFeed.Title,
		Mdate:  time.Now(),
		Cdate:  time.Now(),
	}

	if hub, ok := rawFeed.Custom["hub"]; ok {
		feed.HubURL = &hub
	}

	hasher := md5.New()
	hasher.Write([]byte(url))
	feed.ID = hex.EncodeToString(hasher.Sum(nil))
	return feed, nil
}
