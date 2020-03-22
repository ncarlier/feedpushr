package pipeline

import (
	"fmt"
	"sync"
	"time"

	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/filter"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
	"github.com/ncarlier/feedpushr/v2/pkg/output"
	"github.com/ncarlier/feedpushr/v2/pkg/plugin"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Pipeline is the object that process filters and outputs.
type Pipeline struct {
	lock           sync.RWMutex
	plugins        map[string]model.OutputPlugin
	outputs        []model.Output
	db             store.DB
	ChainFilter    *filter.Chain
	cacheRetention time.Duration
	log            zerolog.Logger
}

// NewPipeline creates a new pipeline
func NewPipeline(db store.DB, cacheRetention time.Duration) (*Pipeline, error) {
	pipeline := &Pipeline{
		plugins:        output.GetBuiltinOutputPlugins(),
		outputs:        []model.Output{},
		db:             db,
		cacheRetention: cacheRetention,
		log:            log.With().Str("component", "pipeline").Logger(),
	}

	// Register external output plugins...
	err := plugin.GetRegistry().ForEachOutputPlugin(func(plug model.OutputPlugin) error {
		pipeline.plugins[plug.Spec().Name] = plug
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Load output outputs from DB
	err = db.ForEachOutput(func(o *model.OutputDef) error {
		if o == nil {
			return fmt.Errorf("output is null")
		}
		_, err := pipeline.AddOutput(o)
		return err
	})
	return pipeline, err
}

// GetAvailableOutputs get all available outputs
func (p *Pipeline) GetAvailableOutputs() []model.Spec {
	result := []model.Spec{}
	for _, plugin := range p.plugins {
		result = append(result, plugin.Spec())
	}
	return result
}

// Process articles
func (p *Pipeline) Process(articles []*model.Article) uint64 {
	var nbProcessedArticles uint64
	maxAge := time.Now().Add(-p.cacheRetention)
	p.log.Debug().Int("items", len(articles)).Str("before", maxAge.String()).Msg("processing articles")
	for _, article := range articles {
		logger := p.log.With().Str("GUID", article.GUID).Logger()
		if err := article.IsValid(maxAge); err != nil {
			logger.Debug().Err(err).Msg("unable to push article")
			continue
		}

		// Send article to all outputs...
		filteredOnce := false
		sentOnce := false
		for _, provider := range p.outputs {
			if !provider.GetDef().Enabled {
				// Ignore output that are disabled
				// TODO also ignore if the article doesn't match the condition
				continue
			}
			logger = logger.With().Str("output", provider.GetDef().Name).Logger()
			// Check that the article is not already sent
			cached, err := p.hasAlreadySent(article, provider)
			if err != nil {
				logger.Debug().Err(err).Msg("unable to get article from cache: ignore")
			}
			if cached {
				logger.Debug().Msg("article already sent")
				continue
			}

			if p.ChainFilter != nil && !filteredOnce {
				// Apply filter chain on article
				err = p.ChainFilter.Apply(article)
				if err != nil {
					logger.Error().Err(err).Msg("unable to apply chain filter on article")
					break
				}
				filteredOnce = true
			}

			// Send article
			err = p.send(article, provider)
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

func (p *Pipeline) hasAlreadySent(article *model.Article, output model.Output) (bool, error) {
	key := common.Hash(article.Hash(), output.GetDef().Hash())
	item, err := p.db.GetFromCache(key)
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

func (p *Pipeline) send(article *model.Article, output model.Output) error {
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
	err = p.db.StoreToCache(key, item)
	if err != nil {
		p.log.Error().Err(err).Str(
			"GUID", article.GUID,
		).Str(
			"provider", output.GetDef().Name,
		).Msg("unable to store article into the cache: ignore")
	}
	return nil
}
