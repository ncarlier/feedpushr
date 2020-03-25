package test

import (
	"testing"
	"time"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v2/pkg/output"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
)

var (
	db      store.DB
	srv     = goa.New("ctrl-test")
	outputs *output.Manager
	aggreg  *aggregator.Manager
)

func setup(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	var err error
	db, err = store.Configure("memory://")
	if err != nil {
		t.Fatalf("Unable to setup Database: %v", err)
	}
	// Init the pipeline
	outputs, err = output.NewManager(db, time.Hour)
	if err != nil {
		t.Fatalf("Unable to setup Output Manager: %v", err)
	}

	aggreg = aggregator.NewManager(outputs, time.Minute, time.Second*5, "")

	return func(t *testing.T) {
		t.Log("teardown test case")
		aggreg.Shutdown()
		defer db.Close()
	}
}
