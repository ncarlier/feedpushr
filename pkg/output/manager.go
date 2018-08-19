package output

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/ncarlier/feedpushr/pkg/filter"
	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/feedpushr/pkg/plugin"
	"github.com/ncarlier/feedpushr/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// GetArticleKey computes the key of a Article.
func getArticleKey(article *model.Article) string {
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
	provider       model.OutputProvider
	db             store.DB
	chainFilter    *filter.Chain
	cacheRetention time.Duration
	log            zerolog.Logger
}

// NewManager creates a new manager
func NewManager(db store.DB, uri string, cacheRetention time.Duration, pr *plugin.Registry, cf *filter.Chain) (*Manager, error) {
	provider, err := newOutputProvider(uri, pr)
	if err != nil {
		return nil, err
	}
	manager := &Manager{
		provider:       provider,
		db:             db,
		chainFilter:    cf,
		cacheRetention: cacheRetention,
		log:            log.With().Str("component", "output").Logger(),
	}
	return manager, nil
}

// Send feeds to the output provider
func (m *Manager) Send(articles []*model.Article) error {
	maxAge := time.Now().Add(-m.cacheRetention)
	m.log.Debug().Int("items", len(articles)).Str("before", maxAge.String()).Msg("processing articles")
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
			// m.log.Debug().Str("GUID", article.GUID).Str("date", (*date).String()).Msg("unable to push article: article too old")
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

		// Apply filter chain on article
		err = m.chainFilter.Apply(article)
		if err != nil {
			m.log.Error().Err(err).Str("GUID", article.GUID).Msg("unable to apply chain filter on article")
			continue
		}

		// Send article...
		err = m.provider.Send(article)
		if err != nil {
			m.log.Error().Err(err).Str("GUID", article.GUID).Msg("unable to push article")
			continue
		}
		m.log.Info().Str("GUID", article.GUID).Msg("article pushed")
		// Set article as sent by updating the cache
		item = &model.CacheItem{
			Value: article.GUID,
			Date:  *date,
		}
		err = m.db.StoreToCache(key, item)
		if err != nil {
			m.log.Error().Err(err).Str("GUID", article.GUID).Msg("unable to store article into the cache")
		}
	}
	return nil
}

// GetSpec return specification of the chain filter
func (m *Manager) GetSpec() model.OutputSpec {
	return m.provider.GetSpec()
}
