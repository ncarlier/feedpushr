package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/autogen/app/test"
	"github.com/ncarlier/feedpushr/v3/pkg/controller"
)

func TestFeedCRUD(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	ctrl := controller.NewFeedController(srv, db, aggregators)
	ctx := context.Background()

	url := "http://rss.cnn.com/rss/edition.rss"

	// CREATE
	status := true
	_, f := test.CreateFeedCreated(t, ctx, srv, ctrl, &status, nil, nil, url)
	assert.Equal(t, "running", *f.Status)
	assert.Equal(t, url, f.XMLURL)
	id := f.ID

	// GET
	_, f = test.GetFeedOK(t, ctx, srv, ctrl, id)
	assert.Equal(t, id, f.ID)
	assert.Equal(t, url, f.XMLURL)

	// FIND
	_, page := test.ListFeedOK(t, ctx, srv, ctrl, 1, nil, 5)
	assert.Equal(t, 5, page.Size)
	assert.Equal(t, 1, page.Current)
	assert.Equal(t, 1, page.Total)
	item := page.Data[0]
	assert.Equal(t, id, item.ID)

	// SEARCH
	q := "international"
	_, page = test.ListFeedOK(t, ctx, srv, ctrl, 1, &q, 5)
	assert.Equal(t, 5, page.Size)
	assert.Equal(t, 1, page.Current)
	assert.Equal(t, 1, page.Total)
	item = page.Data[0]
	assert.Equal(t, id, item.ID)

	// DELETE
	test.DeleteFeedNoContent(t, ctx, srv, ctrl, id)

	// GET 404
	test.GetFeedNotFound(t, ctx, srv, ctrl, id)
}
