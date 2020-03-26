package output

import (
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v2/pkg/cache"
	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/filter"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Processor is the object that operate an output.
type Processor struct {
	output              model.Output
	Filters             *filter.Chain
	cache               *cache.Manager
	nbProcessedArticles uint64
	log                 zerolog.Logger
}

// NewOutputProcessor creates a new output processor
func NewOutputProcessor(out model.Output, filters *filter.Chain, cm *cache.Manager) *Processor {
	return &Processor{
		output:  out,
		Filters: filters,
		cache:   cm,
		log:     log.With().Str("component", "output-processor").Logger(),
	}
}

// Shutdown a processor
func (p *Processor) Shutdown() error {
	// TODO processor gracefull shutdown
	return nil
}

// GetDef returns processor output definition
func (p *Processor) GetDef() model.OutputDef {
	def := p.output.GetDef()
	def.Filters = p.Filters.GetFilterDefs()
	return def
}

// Process articles
func (p *Processor) Process(articles []*model.Article) {
	if !p.output.GetDef().Enabled {
		// Ignore output that are disabled
		return
	}

	maxAge := p.cache.MaxAge()
	p.log.Debug().Int("items", len(articles)).Str("before", maxAge.String()).Msg("processing articles")
	for _, article := range articles {
		logger := p.log.With().Str("GUID", article.GUID).Logger()
		if err := article.IsValid(maxAge); err != nil {
			logger.Debug().Err(err).Msg("unable to push article")
			continue
		}
		// Check that the article is not already sent
		cached, err := p.hasAlreadySent(article)
		if err != nil {
			logger.Debug().Err(err).Msg("unable to get article from cache: ignore")
		}
		if cached {
			logger.Debug().Msg("article already sent")
			continue
		}

		// Apply filters on article
		err = p.Filters.Apply(article)
		if err != nil {
			logger.Error().Err(err).Msg("unable to apply filters on article")
			break
		}

		// Push article
		err = p.push(article)
		if err != nil {
			logger.Error().Err(err).Msg("unable to push article")
			continue
		}
		logger.Info().Msg("article pushed")
		atomic.AddUint64(&p.nbProcessedArticles, 1)
	}
}

func (p *Processor) hasAlreadySent(article *model.Article) (bool, error) {
	key := common.Hash(article.Hash(), p.output.GetDef().Hash())
	item, err := p.cache.Get(key)
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

func (p *Processor) push(article *model.Article) error {
	// Send the article to the output
	err := p.output.Send(article)
	if err != nil {
		return err
	}

	// Set article as sent by updating the cache
	key := common.Hash(article.Hash(), p.output.GetDef().Hash())
	item := &model.CacheItem{
		Value: article.GUID,
		Date:  *article.RefDate(),
	}
	err = p.cache.Set(key, item)
	if err != nil {
		p.log.Error().Err(err).Str(
			"GUID", article.GUID,
		).Str(
			"provider", p.output.GetDef().Name,
		).Msg("unable to store article into the cache: ignore")
	}
	return nil
}
