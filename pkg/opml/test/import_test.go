package test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/helper"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/opml"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
)

var db store.DB

var importer *opml.Importer

var testCases = []struct {
	url string
	tag string
}{
	{"http://www.nofrag.com/nofrag.rss", "games"},
	{"https://www.hashicorp.com/blog/feed.xml", "computer_science"},
}

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	var err error
	db, err = store.NewDB("memory://", model.Quota{})
	if err != nil {
		t.Fatalf("Unable to setup Database: %v", err)
	}
	importer = opml.NewOPMLImporter(db)
	return func(t *testing.T) {
		t.Log("teardown test case")
		defer db.Close()
	}
}

func TestImportSimpleOPML(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	job, err := importer.ImportOPMLFile("./tc_simple.xml")
	assert.Nil(t, err)
	assert.True(t, job.ID > 0, "invalid job ID")
	over := job.Wait(10 * time.Second)
	assert.True(t, over, "job is not over")
	output, err := importer.Get(uint(job.ID))
	assert.Nil(t, err)
	for line := range output {
		assert.True(t, strings.HasSuffix(line, "ok") || line == "done", "invalid job output content")
	}

	assert.True(t, db.ExistsFeed("https://www.hashicorp.com/blog/feed.xml"), "feed should be created")
}

func testImportOPML(t *testing.T, filename string) {
	job, err := importer.ImportOPMLFile(filename)
	assert.Nil(t, err)
	assert.True(t, job.ID > 0, "invalid job ID")
	output, err := importer.Get(uint(job.ID))
	assert.Nil(t, err)
	for line := range output {
		assert.True(t, strings.HasSuffix(line, "ok") || line == "done", "invalid job output content")
	}
	for idx, tc := range testCases {
		assert.True(t, db.ExistsFeed(tc.url), fmt.Sprintf("feed #%d should be created", idx))
		id := helper.Hash(tc.url)
		feed, err := db.GetFeed(id)
		assert.Nil(t, err, fmt.Sprintf("error #%d should be nil", idx))
		assert.NotNil(t, feed, fmt.Sprintf("feed #%d should not be nil", idx))
		assert.Contains(t, feed.Tags, tc.tag, fmt.Sprintf("invalid tags for feed #%d", idx))
	}
}

func TestImportOPMLWithOutlineCategories(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	testImportOPML(t, "./tc_with_outline_categories.xml")
}

func TestImportOPMLWithInlineCategories(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	testImportOPML(t, "./tc_with_inline_categories.xml")
}
