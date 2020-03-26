package test

import (
	"testing"
	"time"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v2/pkg/cache"
	"github.com/ncarlier/feedpushr/v2/pkg/config"
	"github.com/ncarlier/feedpushr/v2/pkg/output"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
)

var (
	db          store.DB
	srv         = goa.New("ctrl-test")
	outputs     *output.Manager
	aggregators *aggregator.Manager
)

func setup(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	var err error
	db, err = store.NewDB("memory://")
	if err != nil {
		t.Fatalf("Unable to setup database: %v", err)
	}
	cm, err := cache.NewCacheManager(db, config.Config{CacheRetention: time.Hour})
	if err != nil {
		t.Fatalf("Unable to setup cache manager: %v", err)
	}

	// Init outputs manager
	outputs, err = output.NewOutputManager(cm)
	if err != nil {
		t.Fatalf("Unable to setup output manager: %v", err)
	}

	aggregators = aggregator.NewAggregatorManager(outputs, time.Minute, time.Second*5, "")

	return func(t *testing.T) {
		t.Log("teardown test case")
		aggregators.Shutdown()
		defer db.Close()
	}
}
