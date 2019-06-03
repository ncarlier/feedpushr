package job

import (
	"time"

	"github.com/ncarlier/feedpushr/pkg/config"
	"github.com/ncarlier/feedpushr/pkg/store"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// CleanCacheJob is a job to clean the cache
type CleanCacheJob struct {
	db     store.DB
	ticker *time.Ticker
	logger zerolog.Logger
}

// NewCleanCacheJob create and start new job to clean the cache
func NewCleanCacheJob(_db store.DB) *CleanCacheJob {
	job := &CleanCacheJob{
		db:     _db,
		ticker: time.NewTicker(24 * time.Hour),
		logger: log.With().Str("job", "cache-buster").Logger(),
	}
	go job.start()
	return job
}

func (ccj *CleanCacheJob) start() {
	ccj.logger.Debug().Str("retention", (*config.CacheRetention).String()).Msg("job started")
	for range ccj.ticker.C {
		ccj.logger.Debug().Msg("running job...")
		maxAge := time.Now().Add(-*config.CacheRetention)
		err := ccj.db.EvictFromCache(maxAge)
		if err != nil {
			ccj.logger.Error().Err(err).Msg("unable to clean cache")
			break
		}
		ccj.logger.Debug().Msg("done")
	}
}

// Stop job
func (ccj *CleanCacheJob) Stop() {
	ccj.ticker.Stop()
	ccj.logger.Debug().Msg("job stopped")
}
