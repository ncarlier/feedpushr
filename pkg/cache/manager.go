package cache

import (
	"time"

	"github.com/ncarlier/feedpushr/v2/pkg/config"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Manager to operate cache
type Manager struct {
	repository store.CacheRepository
	retention  time.Duration
	logger     zerolog.Logger
	ticker     *time.Ticker
}

// NewCacheManager creates nes cache manager
func NewCacheManager(repository store.CacheRepository, conf config.Config) (*Manager, error) {
	logger := log.With().Str("component", "cache-manager").Logger()
	// Clear cache if asked
	if conf.ClearCache {
		logger.Debug().Msg("clearing the cache...")
		if err := repository.ClearCache(); err != nil {
			logger.Error().Err(err).Msg("unable to clear the cache")
			return nil, err
		}
		logger.Info().Msg("cache cleared")
	}

	// Create manager
	manager := &Manager{
		repository: repository,
		retention:  conf.CacheRetention,
		logger:     logger,
		ticker:     time.NewTicker(24 * time.Hour),
	}

	// Start cache buster job
	go manager.startCacheBusterJob()

	return manager, nil
}

// MaxAge compute current cache max age
func (m *Manager) MaxAge() time.Time {
	return time.Now().Add(-m.retention)
}

// Get item from cache
func (m *Manager) Get(key string) (*model.CacheItem, error) {
	return m.repository.GetFromCache(key)
}

// Set item into the cache
func (m *Manager) Set(key string, item *model.CacheItem) error {
	return m.repository.StoreToCache(key, item)
}

func (m *Manager) startCacheBusterJob() {
	m.logger.Debug().Str("retention", m.retention.String()).Msg("cache buster started")
	for range m.ticker.C {
		m.logger.Debug().Msg("cleaning cache...")
		maxAge := time.Now().Add(-m.retention)
		err := m.repository.EvictFromCache(maxAge)
		if err != nil {
			m.logger.Error().Err(err).Msg("unable to clean cache")
			break
		}
		m.logger.Debug().Msg("cache cleaned")
	}
}

// Shutdown cache manager
func (m *Manager) Shutdown() {
	m.ticker.Stop()
	m.logger.Debug().Msg("cache buster stopped")
}
