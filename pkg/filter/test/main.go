
package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/store"
)

var db store.DB

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	var err error
	db, err = store.Configure("memory://")
	if err != nil {
		t.Fatalf("Unable to setup Database: %v", err)
	}
	return func(t *testing.T) {
		t.Log("teardown test case")
		defer db.Close()
	}
}
