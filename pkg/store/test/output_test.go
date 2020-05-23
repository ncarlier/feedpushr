package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

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
	assert.Nil(t, err)

	outputs, err := db.ListOutputs(1, 10)
	assert.Nil(t, err)
	assert.NotNil(t, outputs)
	assert.Len(t, *outputs, 1, "unexpected number of outputs")
	assert.Equal(t, id, (*outputs)[0].ID, "unexpected output ID")

	def, err = db.GetOutput(id)
	assert.Nil(t, err)
	assert.NotNil(t, def)
	assert.Equal(t, id, def.ID, "unexpected output ID")
	assert.Len(t, def.Filters, 1, "unexpected output filters")
	assert.Equal(t, "test", def.Filters[0].Alias, "unexpected output filters")

	_, err = db.DeleteOutput(id)
	assert.Nil(t, err)

	_, err = db.GetOutput(id)
	assert.NotNil(t, err)
	assert.Equal(t, common.ErrOutputNotFound.Error(), err.Error(), "unexpected error message")
}
