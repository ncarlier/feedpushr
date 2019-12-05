package aggregator

import (
	"sync"
	"time"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/output"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Manager of the feed aggregators
type Manager struct {
	feedAggregators   map[string]*FeedAggregator
	shutdownWaitGroup sync.WaitGroup
	output            *output.Manager
	log               zerolog.Logger
	delay             time.Duration
	timeout           time.Duration
	callbackURL       string
}

// NewManager creates a new aggregator manager
func NewManager(om *output.Manager, delay time.Duration, timeout time.Duration, callbackURL string) *Manager {
	return &Manager{
		feedAggregators: make(map[string]*FeedAggregator),
		output:          om,
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
func (m *Manager) RegisterFeedAggregator(feed *app.Feed) *FeedAggregator {
	fa := m.GetFeedAggregator(feed.ID)
	if fa != nil {
		m.log.Debug().Str("source", feed.ID).Msg("feed aggregator already registered")
		return fa
	}
	fa = NewFeedAggregator(feed, m.output, m.delay, m.timeout, m.callbackURL)
	m.feedAggregators[feed.ID] = fa
	m.shutdownWaitGroup.Add(1)
	fa.Start()
	m.log.Debug().Str("source", feed.ID).Msg("feed aggregator registered")
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
	for _, fa := range m.feedAggregators {
		go m.UnRegisterFeedAggregator(fa.id)
	}
	m.shutdownWaitGroup.Wait()
	m.log.Debug().Msg("all aggregators stopped")
}
