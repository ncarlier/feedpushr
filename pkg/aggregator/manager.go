package aggregator

import (
	"sync"
	"time"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/output"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Manager of the feed aggregators
type Manager struct {
	feedAggregators   sync.Map
	shutdownWaitGroup sync.WaitGroup
	outputs           *output.Manager
	log               zerolog.Logger
	delay             time.Duration
	timeout           time.Duration
	callbackURL       string
}

// NewAggregatorManager creates a new aggregator manager
func NewAggregatorManager(outputs *output.Manager, delay time.Duration, timeout time.Duration, callbackURL string) *Manager {
	return &Manager{
		feedAggregators: sync.Map{},
		outputs:         outputs,
		log:             log.With().Str("component", "aggregator").Logger(),
		delay:           delay,
		timeout:         timeout,
		callbackURL:     callbackURL,
	}
}

// GetFeedAggregator returns a feed aggregator
func (m *Manager) GetFeedAggregator(id string) *FeedAggregator {
	fa, found := m.feedAggregators.Load(id)
	if found {
		return fa.(*FeedAggregator)
	} else {
		m.log.Debug().Str("source", id).Msg("feed aggregator not found")
		return nil
	}
}

// RegisterFeedAggregator register and start a new feed aggregator
func (m *Manager) RegisterFeedAggregator(feed *model.FeedDef, delay time.Duration) *FeedAggregator {
	fa := m.GetFeedAggregator(feed.ID)
	if fa != nil {
		m.log.Debug().Str("source", feed.ID).Msg("feed aggregator already registered")
		return fa
	}
	fa = NewFeedAggregator(feed, m.outputs, m.delay, m.timeout, m.callbackURL)
	m.feedAggregators.Store(feed.ID, fa)
	m.shutdownWaitGroup.Add(1)
	if delay > 0 {
		fa.StartWithDelay(delay)
	} else {
		fa.Start()
	}
	m.log.Debug().Str("source", feed.ID).Dur("delay", delay).Msg("feed aggregator registered")
	return fa
}

// UnRegisterFeedAggregator stop and un-register a feed aggregator
func (m *Manager) UnRegisterFeedAggregator(id string) {
	fa := m.GetFeedAggregator(id)
	if fa == nil {
		m.log.Warn().Str("feed", id).Msg("unable to deregister feed aggregator: not found")
		return
	}
	fa.Stop()
	m.feedAggregators.Delete(id)
	m.shutdownWaitGroup.Done()
	m.log.Debug().Str("feed", id).Msg("feed aggregator unregistered")
}

// RestartFeedAggregator restart feed aggregator with delay
func (m *Manager) RestartFeedAggregator(id string, delay time.Duration) {
	fa := m.GetFeedAggregator(id)
	if fa == nil {
		m.log.Warn().Str("feed", id).Msg("unable to restart feed aggregator: not found")
		return
	}
	fa.Stop()
	fa.StartWithDelay(delay)
}

// Shutdown stop the manager (aka. stop and unregister all feed aggregator)
func (m *Manager) Shutdown() {
	m.log.Debug().Msg("shutting down all aggregators")
	// Build temporary list of IDs
	// This is necessary because feddAggregators will be mutated
	var ids []string
	m.feedAggregators.Range(func(key any, value any) bool {
		id := key.(string)
		ids = append(ids, id)
		return true
	})
	for _, id := range ids {
		go m.UnRegisterFeedAggregator(id)
	}
	m.shutdownWaitGroup.Wait()
	m.log.Debug().Msg("all aggregators stopped")
}
