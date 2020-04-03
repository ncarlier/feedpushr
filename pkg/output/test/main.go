package test

import (
	"testing"
	"time"

	"github.com/ncarlier/feedpushr/v2/pkg/cache"
	"github.com/ncarlier/feedpushr/v2/pkg/config"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
)

var (
	db store.DB
	cm *cache.Manager
)

func setup(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	var err error
	db, err = store.NewDB("memory://")
	if err != nil {
		t.Fatalf("Unable to setup database: %v", err)
	}
	cm, err = cache.NewCacheManager(db, config.Config{CacheRetention: time.Hour})
	if err != nil {
		t.Fatalf("Unable to setup cache manager: %v", err)
	}

	return func(t *testing.T) {
		t.Log("teardown test case")
		cm.Shutdown()
		defer db.Close()
	}
}
