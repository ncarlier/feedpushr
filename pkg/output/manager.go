package output

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/ncarlier/feedpushr/pkg/cache"
	"github.com/ncarlier/feedpushr/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// GetArticleKey computes the key of a GoFeed item.
func getArticleKey(article *gofeed.Item) string {
	key := article.GUID
	if key == "" {
		key = article.Link
	}
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Manager of output channel.
type Manager struct {
	provider       Provider
	db             store.DB
	cacheRetention time.Duration
	log            zerolog.Logger
}

// NewManager creates a new manager
func NewManager(db store.DB, uri string, cacheRetention time.Duration) (*Manager, error) {
	provider, err := newOutputProvider(uri)
	if err != nil {
		return nil, err
	}
	manager := &Manager{
		provider: provider,
		db:       db,
		log:      log.With().Str("component", "output").Logger(),
	}
	return manager, nil
}

// Send feeds to the output provider
func (m *Manager) Send(articles []*gofeed.Item) error {
	m.log.Debug().Int("items", len(articles)).Msg("processing articles")
	maxAge := time.Now().Add(-m.cacheRetention)
	for _, article := range articles {
		// Ignore old articles
		var date *time.Time
		if article.PublishedParsed != nil {
			date = article.PublishedParsed
		}
		if article.UpdatedParsed != nil {
			date = article.UpdatedParsed
		}
		if date == nil {
			m.log.Debug().Msg("unable to push article: missing article date")
			continue
		}
		if date.Before(maxAge) {
			// Article too old: ignore
			// m.log.Debug().Msg("unable to push article: article too old")
			continue
		}
		// Ignore already sent articles
		key := getArticleKey(article)
		item, err := m.db.GetFromCache(key)
		if err != nil {
			m.log.Debug().Err(err).Msg("unable to retrieve article from cache: sending")
		} else if item != nil {
			if date.After(item.Date) {
				m.log.Debug().Msg("article updated since last push: re-sending")
			} else {
				// Article already sent: ignore
				continue
			}
		}
		// Send article...
		err = m.provider.Send(article)
		if err != nil {
			m.log.Error().Err(err).Str("guid", article.GUID).Msg("unable to send article")
			continue
		}
		// Set article as sent by updating the cache
		item = &cache.Item{
			Value: article.GUID,
			Date:  *date,
		}
		err = m.db.StoreToCache(key, item)
		if err != nil {
			m.log.Error().Err(err).Str("guid", article.GUID).Msg("unable to store article into the cache")
		}
	}
	return nil
}
