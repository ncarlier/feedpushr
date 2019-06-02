package aggregator

import (
	"sync/atomic"
	"time"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/output"
	"github.com/ncarlier/feedpushr/pkg/pshb"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// FeedAggregator aggregator
type FeedAggregator struct {
	id                string
	managementChannel chan Action
	shutdownChannel   chan Action
	stopChannel       chan bool
	log               zerolog.Logger
	feed              *app.Feed
	handler           *FeedHandler
	outputsManager    *output.Manager
	status            Status
	delay             time.Duration
	nextCheck         time.Time
	lastCheck         time.Time
	nbProcessedItems  uint64
	callbackURL       string
}

// NewFeedAggregator creats a new feed aggregator
func NewFeedAggregator(feed *app.Feed, om *output.Manager, delay time.Duration, timeout time.Duration, callbackURL string) *FeedAggregator {
	handler := NewFeedHandler(feed, timeout)
	aggregator := FeedAggregator{
		id:                feed.ID,
		managementChannel: make(chan Action),
		shutdownChannel:   make(chan Action),
		stopChannel:       make(chan bool),
		feed:              feed,
		handler:           handler,
		outputsManager:    om,
		status:            StoppedStatus,
		delay:             delay,
		log:               log.With().Str("aggregator", feed.ID).Logger(),
		callbackURL:       callbackURL,
	}
	aggregator.run()
	return &aggregator
}

func (fa *FeedAggregator) running() {
	if fa.status == RunningStatus {
		return
	}
	go func() {
		fa.status = RunningStatus
		for {
			delay := time.Until(fa.nextCheck)
			if delay.Seconds() <= 0 {
				fa.log.Debug().Str("status", fa.status.String()).Msg("fetching feed")
				// Refresh the feed
				status, items := fa.handler.Refresh()
				if items != nil && len(items) > 0 {
					// Send feed's articles to the output provider
					nb := fa.outputsManager.Send(items)
					atomic.AddUint64(&fa.nbProcessedItems, nb)
				}
				if fa.feed.HubURL != nil && *fa.feed.HubURL != "" && fa.callbackURL != "" {
					// Send subscription request if the feed is bound to a hub
					err := pshb.Subscribe(*fa.feed.HubURL, fa.feed.XMLURL, fa.callbackURL)
					if err != nil {
						fa.log.Error().Err(err).Str("hub", *fa.feed.HubURL).Msg("unable to subscribe to the Hub")
					} else {
						fa.log.Debug().Str("hub", *fa.feed.HubURL).Msg("subscribed to the Hub")
					}
				}
				fa.lastCheck = time.Now()
				fa.nextCheck = status.ComputeNextCheckDate(fa.delay)
				delay = time.Until(fa.nextCheck)
			}
			fa.log.Debug().Str("duration", delay.String()).Msg("waiting")
			select {
			case <-time.After(delay):
			case <-fa.stopChannel:
				fa.status = StoppedStatus
				fa.shutdownChannel <- StopAction
				return
			}
		}
	}()
}

func (fa *FeedAggregator) run() {
	go func() {
		for {
			select {
			case action := <-fa.managementChannel:
				// Receive a management request.
				switch action {
				case StartAction:
					fa.log.Debug().Str("status", fa.status.String()).Msg("starting feed aggregator...")
					if fa.status == StoppedStatus {
						fa.running()
					}
				case StopAction:
					fa.log.Debug().Str("status", fa.status.String()).Msg("stopping feed aggregator...")
					if fa.status == RunningStatus {
						fa.stopChannel <- true
					} else {
						fa.shutdownChannel <- StopAction
					}
				default:
					fa.log.Warn().Str("action", action.String()).Msg("unable to handle action")
				}
			}
		}
	}()
}

// Stop stops feed aggregator.
func (fa *FeedAggregator) Stop() {
	fa.managementChannel <- StopAction
	<-fa.shutdownChannel
	fa.log.Debug().Msg("feed aggregator stopped")
}

// Start starts feed aggregator.
func (fa *FeedAggregator) Start() {
	fa.managementChannel <- StartAction
}

// StartWithDelay starts feed aggregator with a delay.
func (fa *FeedAggregator) StartWithDelay(delay time.Duration) {
	fa.nextCheck = time.Now().Add(delay)
	fa.Start()
}

// GetFeedWithAggregationStatus get a copy of the aggregator feed hydrated with aggregation status.
func (fa *FeedAggregator) GetFeedWithAggregationStatus() *app.Feed {
	result := *fa.feed
	status := fa.status.String()
	result.Status = &status
	lastCheck := fa.lastCheck
	result.LastCheck = &lastCheck
	nextCheck := fa.nextCheck
	result.NextCheck = &nextCheck
	nbProcessedItems := int(fa.nbProcessedItems)
	result.NbProcessedItems = &nbProcessedItems
	if fa.handler.status.ErrorMsg != "" {
		msg := fa.handler.status.ErrorMsg
		result.ErrorMsg = &msg
	}
	if fa.handler.status.ErrorCount != 0 {
		count := fa.handler.status.ErrorCount
		result.ErrorCount = &count
	}
	if fa.callbackURL != "" && fa.feed.HubURL != nil {
		result.HubURL = pshb.GetSubscriptionDetailsURL(*fa.feed.HubURL, fa.feed.XMLURL, fa.callbackURL)
	} else {
		result.HubURL = nil
	}

	return &result
}
