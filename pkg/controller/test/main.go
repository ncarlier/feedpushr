package test

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v3/pkg/cache"
	"github.com/ncarlier/feedpushr/v3/pkg/config"
	"github.com/ncarlier/feedpushr/v3/pkg/filter"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/output"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
)

var (
	db          store.DB
	srv         = goa.New("ctrl-test")
	chain       *filter.Chain
	outputs     *output.Manager
	aggregators *aggregator.Manager
)

func setup(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")

	workspace, err := ioutil.TempDir("", "feedpushr")
	if err != nil {
		t.Fatalf("unable to setup temporary workspace: %v", err)
	}

	db, err = store.NewDB("boltdb://"+workspace+"/test.db", model.Quota{})
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

	// Init empty chain filter
	chain, err = filter.NewChainFilter(model.FilterDefCollection{})
	if err != nil {
		t.Fatalf("Unable to setup chain filter: %v", err)
	}

	aggregators = aggregator.NewAggregatorManager(outputs, time.Minute, time.Second*5, "")

	return func(t *testing.T) {
		t.Log("teardown test case")
		aggregators.Shutdown()
		outputs.Shutdown()
		cm.Shutdown()
		defer db.Close()
		defer os.RemoveAll(workspace)
	}
}
