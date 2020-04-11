package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ncarlier/feedpushr/v3/pkg/assert"
	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

func TestOutputCRUD(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	id := uuid.New().String()
	def := &model.OutputDef{
		ID: id,
		Filters: model.FilterDefCollection{
			&model.FilterDef{
				Alias: "test",
			},
		},
	}
	_, err := db.SaveOutput(*def)
	assert.Nil(t, err, "should be nil")

	outputs, err := db.ListOutputs(1, 10)
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, outputs, "should not be nil")
	assert.Equal(t, 1, len(*outputs), "unexpected number of outputs")
	assert.Equal(t, id, (*outputs)[0].ID, "unexpected output ID")

	def, err = db.GetOutput(id)
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, def, "should not be nil")
	assert.Equal(t, id, def.ID, "unexpected output ID")
	assert.Equal(t, 1, len(def.Filters), "unexpected output filters")
	assert.Equal(t, "test", def.Filters[0].Alias, "unexpected output filters")

	_, err = db.DeleteOutput(id)
	assert.Nil(t, err, "should be nil")

	_, err = db.GetOutput(id)
	assert.NotNil(t, err, "should not be nil")
	assert.Equal(t, common.ErrOutputNotFound.Error(), err.Error(), "unexpected error message")
}
