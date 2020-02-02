package service

import (
	"fmt"

	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
)

func loadFeedAggregators(db store.DB, m *aggregator.Manager) error {
	return db.ForEachFeed(func(f *app.Feed) error {
		if f == nil {
			return fmt.Errorf("feed is null")
		}
		// TODO do a progressive load increase
		if f.Status != nil && *f.Status == aggregator.RunningStatus.String() {
			m.RegisterFeedAggregator(f)
		}
		return nil
	})
}
