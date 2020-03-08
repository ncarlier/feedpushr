package aggregator

import (
	"sync"
	"time"

	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/pipeline"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Manager of the feed aggregators
type Manager struct {
	feedAggregators   map[string]*FeedAggregator
	shutdownWaitGroup sync.WaitGroup
	pipeline          *pipeline.Pipeline
	log               zerolog.Logger
	delay             time.Duration
	timeout           time.Duration
	callbackURL       string
}

// NewManager creates a new aggregator manager
func NewManager(pipe *pipeline.Pipeline, delay time.Duration, timeout time.Duration, callbackURL string) *Manager {
	return &Manager{
		feedAggregators: make(map[string]*FeedAggregator),
		pipeline:        pipe,
		log:             log.With().Str("component", "aggregator").Logger(),
		delay:           delay,
		timeout:         timeout,
		callbackURL:     callbackURL,
	}
}

// GetFeedAggregator returns a feed aggregator
func (m *Manager) GetFeedAggregator(id string) *FeedAggregator {
	return m.feedAggregators[id]
}

// RegisterFeedAggregator register and start a new feed aggregator
func (m *Manager) RegisterFeedAggregator(feed *app.Feed, delay time.Duration) *FeedAggregator {
	fa := m.GetFeedAggregator(feed.ID)
	if fa != nil {
		m.log.Debug().Str("source", feed.ID).Msg("feed aggregator already registered")
		return fa
	}
	fa = NewFeedAggregator(feed, m.pipeline, m.delay, m.timeout, m.callbackURL)
	m.feedAggregators[feed.ID] = fa
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
	m.feedAggregators[id] = nil
	delete(m.feedAggregators, id)
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
	// This is necessary because feddAggregators will be mutate
	ids := make([]string, 0, len(m.feedAggregators))
	for id := range m.feedAggregators {
		ids = append(ids, id)
	}
	for _, id := range ids {
		go m.UnRegisterFeedAggregator(id)
	}
	m.shutdownWaitGroup.Wait()
	m.log.Debug().Msg("all aggregators stopped")
}
