package test

import (
	"context"
	"testing"

	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/autogen/app/test"
	"github.com/ncarlier/feedpushr/v2/pkg/assert"
	"github.com/ncarlier/feedpushr/v2/pkg/controller"
)

func TestOutputCRUD(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	ctrl := controller.NewOutputController(srv, db, pipe)
	ctx := context.Background()

	// CREATE
	alias := "test output"
	payload := &app.CreateOutputPayload{
		Alias:     alias,
		Name:      "stdout",
		Condition: "\"test\" in Tags",
	}
	_, out := test.CreateOutputCreated(t, ctx, srv, ctrl, payload)
	assert.Equal(t, alias, out.Alias, "")
	assert.Equal(t, "stdout", out.Name, "")
	assert.Equal(t, false, out.Enabled, "")
	assert.Equal(t, "\"test\" in Tags", out.Condition, "")
	assert.Equal(t, uint64(0), out.Props["nbSuccess"], "")
	id := out.ID

	// GET
	_, out = test.GetOutputOK(t, ctx, srv, ctrl, id)
	assert.Equal(t, id, out.ID, "")
	assert.Equal(t, "test output", out.Alias, "")
	assert.Equal(t, "stdout", out.Name, "")

	// FIND
	_, list := test.ListOutputOK(t, ctx, srv, ctrl)
	assert.True(t, len(list) > 0, "")
	item := list[len(list)-1]
	assert.Equal(t, id, item.ID, "")

	// UPDATE
	alias = "updated output"
	update := &app.UpdateOutputPayload{
		Alias:   &alias,
		Enabled: true,
	}
	_, out = test.UpdateOutputOK(t, ctx, srv, ctrl, id, update)
	assert.Equal(t, id, out.ID, "")
	assert.Equal(t, alias, out.Alias, "")
	assert.Equal(t, "stdout", out.Name, "")
	assert.Equal(t, "\"test\" in Tags", out.Condition, "")
	assert.Equal(t, true, out.Enabled, "")

	// DELETE
	test.DeleteOutputNoContent(t, ctx, srv, ctrl, id)

	// GET 404
	test.GetOutputNotFound(t, ctx, srv, ctrl, id)
}

func TestOutputDefs(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	ctrl := controller.NewOutputController(srv, db, pipe)
	ctx := context.Background()

	_, specs := test.SpecsOutputOK(t, ctx, srv, ctrl)
	assert.True(t, len(specs) > 0, "")
	for _, spec := range specs {
		if spec.Name == "http" {
			assert.Equal(t, 3, len(spec.Props), "")
			assert.Equal(t, "url", spec.Props[0].Name, "")
			assert.Equal(t, "Target URL", spec.Props[0].Desc, "")
			assert.Equal(t, "url", spec.Props[0].Type, "")
		}
	}
}
