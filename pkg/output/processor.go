package output

import (
	"sync"

	"github.com/ncarlier/feedpushr/v2/pkg/cache"
	"github.com/ncarlier/feedpushr/v2/pkg/expr"
	"github.com/ncarlier/feedpushr/v2/pkg/filter"
	"github.com/ncarlier/feedpushr/v2/pkg/helper"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Processor is the object that operate an output.
type Processor struct {
	condition         *expr.ConditionalExpression
	output            model.Output
	Filters           *filter.Chain
	cache             *cache.Manager
	processingChannel chan *model.Article
	stopChannel       chan bool
	stopWaitGroup     sync.WaitGroup
	status            Status
	logger            zerolog.Logger
}

// NewOutputProcessor creates a new output processor
func NewOutputProcessor(out model.Output, filters *filter.Chain, cm *cache.Manager) (*Processor, error) {
	condition, err := expr.NewConditionalExpression(out.GetDef().Condition)
	if err != nil {
		return nil, err
	}
	logger := log.With().Str(
		"component",
		"output-processor",
	).Str(
		"output",
		out.GetDef().ID,
	).Logger()
	processor := &Processor{
		condition:         condition,
		output:            out,
		Filters:           filters,
		cache:             cm,
		status:            StoppedStatus,
		processingChannel: make(chan *model.Article, 1000),
		stopChannel:       make(chan bool),
		logger:            logger,
	}
	processor.start()
	return processor, nil
}

// Update update output and filter chain
func (p *Processor) Update(out model.Output, filters *filter.Chain) {
	p.Shutdown()
	p.output = out
	p.Filters = filters
	p.start()
}

// Shutdown a processor
func (p *Processor) Shutdown() error {
	if p.status == RunningStatus {
		p.stopChannel <- true
		p.stopWaitGroup.Wait()
	}
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
	// Ignore output that are disabled
	if !p.output.GetDef().Enabled {
		return
	}

	p.logger.Debug().Int("items", len(articles)).Msg("processing articles")
	for _, article := range articles {
		if p.condition.Match(article) {
			// Put article into the processing channel
			p.processingChannel <- article
		}
	}
}

func (p *Processor) start() {
	if p.status == RunningStatus {
		return
	}
	go func() {
		p.status = RunningStatus
		p.stopWaitGroup.Add(1)
		for {
			select {
			case article := <-p.processingChannel:
				// TODO error management?
				p.process(article)
			case <-p.stopChannel:
				p.status = StoppedStatus
				p.stopWaitGroup.Done()
				return
			}
		}
	}()
}

func (p *Processor) process(article *model.Article) error {
	logger := p.logger.With().Str("GUID", article.GUID).Logger()

	// Ignore old articles
	maxAge := p.cache.MaxAge()
	if err := article.IsValid(maxAge); err != nil {
		logger.Debug().Err(err).Msg("unable to push article")
		return nil
	}

	// Ignore already sent article
	cached, err := p.hasAlreadySent(article)
	if err != nil {
		logger.Debug().Err(err).Msg("unable to get article from cache: ignore")
	}
	if cached {
		logger.Debug().Msg("article already sent")
		return nil
	}

	// Apply filters on article
	if err := p.Filters.Apply(article); err != nil {
		logger.Error().Err(err).Msg("unable to apply filters on article")
		return err
	}

	// Send the article to the output
	sent, err := p.output.Send(article)
	if err != nil {
		logger.Error().Err(err).Msg("unable to push article")
		return err
	}

	if sent {
		// Set article as sent by updating the cache
		key := helper.Hash(article.Hash(), p.output.GetDef().Hash())
		item := &model.CacheItem{
			Value: article.GUID,
			Date:  *article.RefDate(),
		}
		err = p.cache.Set(key, item)
		if err != nil {
			logger.Error().Err(err).Msg("unable to store article into the cache: ignore")
		}
		logger.Info().Msg("article pushed")
	}
	return nil
}

func (p *Processor) hasAlreadySent(article *model.Article) (bool, error) {
	key := helper.Hash(article.Hash(), p.output.GetDef().Hash())
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
