package opml_test

import (
	"fmt"
	"testing"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/opml"
	"github.com/ncarlier/feedpushr/pkg/store"
)

var db store.DB

var testCases = []struct {
	url string
	tag string
}{
	{"http://www.nofrag.com/nofrag.rss", "games"},
	{"http://www.howtoforge.com/feed.rss", "computer_science"},
}

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

func TestImportSimpleOPML(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	o, err := opml.NewOPMLFromFile("./tc_simple.xml")
	assert.Nil(t, err, "error should be nil")
	err = opml.ImportOPMLToDB(o, db)
	assert.Nil(t, err, "error should be nil")

	assert.True(t, db.ExistsFeed("http://www.hashicorp.com/feed.xml"), "feed should be created")
}

func TestImportOPMLWithCategories(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	o, err := opml.NewOPMLFromFile("./tc_with_categories.xml")
	assert.Nil(t, err, "error should be nil")
	err = opml.ImportOPMLToDB(o, db)
	assert.Nil(t, err, "error should be nil")

	for idx, tc := range testCases {
		assert.True(t, db.ExistsFeed(tc.url), fmt.Sprintf("feed #%d should be created", idx))
		id := common.Hash(tc.url)
		feed, err := db.GetFeed(id)
		assert.Nil(t, err, fmt.Sprintf("error #%d should be nil", idx))
		assert.NotNil(t, feed, fmt.Sprintf("feed #%d should  notbe nil", idx))
		assert.ContainsStr(t, tc.tag, feed.Tags, fmt.Sprintf("invalid tags for feed #%d", idx))
	}
}
