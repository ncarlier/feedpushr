package output

import (
	"fmt"
	"sync"
	"time"

	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/filter"
	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/feedpushr/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Manager of output channel.
type Manager struct {
	lock           sync.RWMutex
	providers      []model.OutputProvider
	db             store.DB
	ChainFilter    *filter.Chain
	cacheRetention time.Duration
	log            zerolog.Logger
}

// NewManager creates a new manager
func NewManager(db store.DB, cacheRetention time.Duration) (*Manager, error) {
	manager := &Manager{
		providers:      []model.OutputProvider{},
		db:             db,
		cacheRetention: cacheRetention,
		log:            log.With().Str("component", "output").Logger(),
	}
	err := db.ForEachOutput(func(o *model.OutputDef) error {
		if o == nil {
			return fmt.Errorf("output is null")
		}
		_, err := manager.Add(o)
		return err
	})
	return manager, err
}

// Send feeds to the output provider
func (m *Manager) Send(articles []*model.Article) uint64 {
	var nbProcessedArticles uint64
	maxAge := time.Now().Add(-m.cacheRetention)
	m.log.Debug().Int("items", len(articles)).Str("before", maxAge.String()).Msg("processing articles")
	for _, article := range articles {
		logger := m.log.With().Str("GUID", article.GUID).Logger()
		if err := article.IsValid(maxAge); err != nil {
			logger.Debug().Err(err).Msg("unable to push article")
			continue
		}

		// Send article to all outputs...
		filteredOnce := false
		sentOnce := false
		for _, provider := range m.providers {
			tags := provider.GetDef().Tags
			if !provider.GetDef().Enabled || !article.Match(tags) {
				// Ignore output that are disabled or don't match article tags
				continue
			}
			logger = logger.With().Str("output", provider.GetDef().Name).Logger()
			// Check that the article is not already sent
			cached, err := m.hasAlreadySent(article, provider)
			if err != nil {
				logger.Debug().Err(err).Msg("unable to get article from cache: ignore")
			}
			if cached {
				logger.Debug().Msg("article already sent")
				continue
			}

			if m.ChainFilter != nil && !filteredOnce {
				// Apply filter chain on article
				err = m.ChainFilter.Apply(article)
				if err != nil {
					logger.Error().Err(err).Msg("unable to apply chain filter on article")
					break
				}
				filteredOnce = true
			}

			// Send article
			err = m.send(article, provider)
			if err != nil {
				logger.Error().Err(err).Msg("unable to push article")
				continue
			}
			sentOnce = true
			logger.Info().Msg("article pushed")
		}
		if sentOnce {
			nbProcessedArticles++
		}
	}
	return nbProcessedArticles
}

func (m *Manager) hasAlreadySent(article *model.Article, output model.OutputProvider) (bool, error) {
	key := common.Hash(article.Hash(), output.GetDef().Hash())
	item, err := m.db.GetFromCache(key)
	if err != nil {
		return false, err
	} else if item != nil {
		date := article.RefDate()
		if date != nil && !date.After(item.Date) {
			// Article already sent
			return true, nil
		}
		// else article updated since last push: re-sending
	}
	return false, nil
}

func (m *Manager) send(article *model.Article, output model.OutputProvider) error {
	// Send the article
	err := output.Send(article)
	if err != nil {
		return err
	}

	// Set article as sent by updating the cache
	key := common.Hash(article.Hash(), output.GetDef().Hash())
	item := &model.CacheItem{
		Value: article.GUID,
		Date:  *article.RefDate(),
	}
	err = m.db.StoreToCache(key, item)
	if err != nil {
		m.log.Error().Err(err).Str(
			"GUID", article.GUID,
		).Str(
			"provider", output.GetDef().Name,
		).Msg("unable to store article into the cache: ignore")
	}
	return nil
}

// GetOutputDefs return all output definitions of the manager
func (m *Manager) GetOutputDefs() []model.OutputDef {
	result := make([]model.OutputDef, len(m.providers))
	for idx, provider := range m.providers {
		result[idx] = provider.GetDef()
	}
	return result
}
