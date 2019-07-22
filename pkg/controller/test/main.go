package test

import (
	"testing"
	"time"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/filter"
	"github.com/ncarlier/feedpushr/pkg/output"
	"github.com/ncarlier/feedpushr/pkg/store"
)

var (
	db  store.DB
	srv = goa.New("ctrl-test")
	om  *output.Manager
)

func setup(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	var err error
	db, err = store.Configure("memory://")
	if err != nil {
		t.Fatalf("Unable to setup Database: %v", err)
	}
	// Init output manager
	om, err = output.NewManager(db, time.Hour)
	if err != nil {
		t.Fatalf("Unable to setup Output Manager: %v", err)
	}
	om.ChainFilter = filter.NewChainFilter()
	if filter, err := builder.NewFilterFromURI("title://?prefix=[test]"); err != nil {
		t.Fatalf("Unable to setup Chain filter: %v", err)
	} else {
		om.ChainFilter.Add(filter)
	}

	return func(t *testing.T) {
		t.Log("teardown test case")
		defer db.Close()
	}
}
