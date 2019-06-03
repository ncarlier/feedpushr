package aggregator

import (
	"fmt"
	"sync"
	"time"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/output"
	"github.com/ncarlier/feedpushr/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Manager of the feed aggregators
type Manager struct {
	feedAggregators   map[string]*FeedAggregator
	shutdownWaitGroup sync.WaitGroup
	db                store.DB
	output            *output.Manager
	log               zerolog.Logger
	delay             time.Duration
	timeout           time.Duration
	callbackURL       string
}

// NewManager creates a new aggregator manager
func NewManager(db store.DB, om *output.Manager, delay time.Duration, timeout time.Duration, callbackURL string) *Manager {
	return &Manager{
		feedAggregators: make(map[string]*FeedAggregator),
		db:              db,
		output:          om,
		log:             log.With().Str("component", "aggregator").Logger(),
		delay:           delay,
		timeout:         timeout,
		callbackURL:     callbackURL,
	}
}

// Start the aggrgator manager (aka. register and start all feed aggregator)
func (m *Manager) Start() error {
	m.log.Debug().Msg("loading feed aggregators...")
	err := m.db.ForEachFeed(func(f *app.Feed) error {
		if f == nil {
			return fmt.Errorf("feed is null")
		}
		// TODO do a progressive load increase
		if f.Status != nil && *f.Status == RunningStatus.String() {
			m.RegisterFeedAggregator(f)
		}
		return nil
	})
	if err != nil {
		return err
	}
	m.log.Info().Int("feeds", len(m.feedAggregators)).Msg("aggregation started")
	return nil
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
		m.log.Warn().Str("feed", id).Msg("unable to unregister feed aggregator: not found")
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
