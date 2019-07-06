package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/common"
)

func TestFilterCRUD(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	filter := &app.Filter{
		ID: 0,
	}
	_, err := db.SaveFilter(filter)
	assert.Nil(t, err, "should be nil")

	filters, err := db.ListFilters(1, 10)
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, filters, "should not be nil")
	assert.Equal(t, 1, len(*filters), "unexpected number of filters")
	assert.Equal(t, 1, (*filters)[0].ID, "unexpected feed ID")
	filter, err = db.GetFilter(1)
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, filter, "should not be nil")
	assert.Equal(t, 1, filter.ID, "unexpected feed ID")
	_, err = db.DeleteFilter(1)
	assert.Nil(t, err, "should be nil")
	_, err = db.GetFilter(1)
	assert.NotNil(t, err, "should not be nil")
	assert.Equal(t, common.ErrFilterNotFound.Error(), err.Error(), "unexpected error message")
}
