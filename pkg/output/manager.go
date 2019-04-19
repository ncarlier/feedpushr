package output

import (
	"time"

	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/filter"
	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/feedpushr/pkg/plugin"
	"github.com/ncarlier/feedpushr/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Manager of output channel.
type Manager struct {
	providers      []model.OutputProvider
	db             store.DB
	chainFilter    *filter.Chain
	cacheRetention time.Duration
	log            zerolog.Logger
}

// NewManager creates a new manager
func NewManager(db store.DB, uris []string, cacheRetention time.Duration, pr *plugin.Registry, cf *filter.Chain) (*Manager, error) {
	providers := []model.OutputProvider{}
	for _, uri := range uris {
		provider, err := newOutputProvider(uri, pr)
		if err != nil {
			return nil, err
		}
		providers = append(providers, provider)
	}
	manager := &Manager{
		providers:      providers,
		db:             db,
		chainFilter:    cf,
		cacheRetention: cacheRetention,
		log:            log.With().Str("component", "output").Logger(),
	}
	return manager, nil
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
			tags := provider.GetSpec().Tags
			if !article.Match(tags) {
				// Ignore output that do not match the article tags
				continue
			}
			logger = logger.With().Str("output", provider.GetSpec().Name).Logger()
			// Check that the article is not already sent
			cached, err := m.hasAlreadySent(article, provider)
			if err != nil {
				logger.Debug().Err(err).Msg("unable to get article from cache: ignore")
			}
			if cached {
				logger.Debug().Msg("article already sent")
				continue
			}

			if !filteredOnce {
				// Apply filter chain on article
				err = m.chainFilter.Apply(article)
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
	key := common.Hash(article.Hash(), output.GetSpec().Hash())
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
	key := common.Hash(article.Hash(), output.GetSpec().Hash())
	item := &model.CacheItem{
		Value: article.GUID,
		Date:  *article.RefDate(),
	}
	err = m.db.StoreToCache(key, item)
	if err != nil {
		m.log.Error().Err(err).Str(
			"GUID", article.GUID,
		).Str(
			"provider", output.GetSpec().Name,
		).Msg("unable to store article into the cache: ignore")
	}
	return nil
}

// GetSpec return specification of the manager
func (m *Manager) GetSpec() []model.OutputSpec {
	result := make([]model.OutputSpec, len(m.providers))
	for idx, provider := range m.providers {
		result[idx] = provider.GetSpec()
	}
	return result
}
