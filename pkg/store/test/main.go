package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
)

var db store.DB

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	var err error
	db, err = store.NewDB("memory://", model.Quota{})
	// db, err = store.NewDB("boltdb:///tmp/test.db", model.Quota{})
	if err != nil {
		t.Fatalf("Unable to setup Database: %v", err)
	}
	return func(t *testing.T) {
		t.Log("teardown test case")
		defer db.Close()
	}
}
