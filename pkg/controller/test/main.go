package test

import (
	"testing"
	"time"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v2/pkg/builder"
	"github.com/ncarlier/feedpushr/v2/pkg/filter"
	"github.com/ncarlier/feedpushr/v2/pkg/pipeline"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
)

var (
	db     store.DB
	srv    = goa.New("ctrl-test")
	pipe   *pipeline.Pipeline
	aggreg *aggregator.Manager
)

func setup(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	var err error
	db, err = store.Configure("memory://")
	if err != nil {
		t.Fatalf("Unable to setup Database: %v", err)
	}
	// Init the pipeline
	pipe, err = pipeline.NewPipeline(db, time.Hour)
	if err != nil {
		t.Fatalf("Unable to setup Output Manager: %v", err)
	}
	pipe.ChainFilter = filter.NewChainFilter()
	filter := builder.NewFilterBuilder().FromURI("title://?prefix=[test]").Build()
	pipe.ChainFilter.Add(filter)

	aggreg = aggregator.NewManager(pipe, time.Minute, time.Second*5, "")

	return func(t *testing.T) {
		t.Log("teardown test case")
		aggreg.Shutdown()
		defer db.Close()
	}
}
