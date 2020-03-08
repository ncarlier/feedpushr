package service

import (
	"fmt"
	"time"

	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
)

func loadFeedAggregators(db store.DB, m *aggregator.Manager, interval time.Duration) error {
	// Delay used to manage progressive load increase
	delay := interval
	return db.ForEachFeed(func(f *app.Feed) error {
		if f == nil {
			return fmt.Errorf("feed is null")
		}
		if f.Status != nil && *f.Status == aggregator.RunningStatus.String() {
			m.RegisterFeedAggregator(f, delay)
			delay += interval
		}
		return nil
	})
}
