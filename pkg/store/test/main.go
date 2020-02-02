package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/v2/pkg/store"
)

var db store.DB

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	var err error
	db, err = store.Configure("memory://")
	// db, err = store.Configure("boltdb:///tmp/test.db")
	if err != nil {
		t.Fatalf("Unable to setup Database: %v", err)
	}
	return func(t *testing.T) {
		t.Log("teardown test case")
		defer db.Close()
	}
}
