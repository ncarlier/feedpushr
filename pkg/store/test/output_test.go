package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/model"
)

func TestOutputCRUD(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	output := &model.OutputDef{
		ID: 0,
	}
	_, err := db.SaveOutput(*output)
	assert.Nil(t, err, "should be nil")

	outputs, err := db.ListOutputs(1, 10)
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, outputs, "should not be nil")
	assert.Equal(t, 1, len(*outputs), "unexpected number of outputs")
	assert.Equal(t, 1, (*outputs)[0].ID, "unexpected feed ID")
	output, err = db.GetOutput(1)
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, output, "should not be nil")
	assert.Equal(t, 1, output.ID, "unexpected feed ID")
	_, err = db.DeleteOutput(1)
	assert.Nil(t, err, "should be nil")
	_, err = db.GetOutput(1)
	assert.NotNil(t, err, "should not be nil")
	assert.Equal(t, common.ErrOutputNotFound.Error(), err.Error(), "unexpected error message")
}
