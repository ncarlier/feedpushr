package test

import (
	"context"
	"testing"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/autogen/app/test"
	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/controller"
)

func TestOutputCRUD(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	ctrl := controller.NewOutputController(srv, db, om)
	ctx := context.Background()

	// CREATE
	tags := "test"
	payload := &app.CreateOutputPayload{
		Name: "stdout",
		Tags: &tags,
	}
	_, out := test.CreateOutputCreated(t, ctx, srv, ctrl, payload)
	assert.Equal(t, "stdout", out.Name, "")
	assert.Equal(t, false, out.Enabled, "")
	assert.ContainsStr(t, "test", out.Tags, "")
	assert.Equal(t, uint64(0), out.Props["nbSuccess"], "")
	id := out.ID

	// GET
	_, out = test.GetOutputOK(t, ctx, srv, ctrl, id)
	assert.Equal(t, id, out.ID, "")
	assert.Equal(t, "stdout", out.Name, "")

	// FIND
	_, list := test.ListOutputOK(t, ctx, srv, ctrl)
	assert.True(t, len(list) > 0, "")
	item := list[len(list)-1]
	assert.Equal(t, id, item.ID, "")

	// UPDATE
	tags = "test,foo"
	update := &app.UpdateOutputPayload{
		Enabled: true,
		Tags:    &tags,
	}
	_, out = test.UpdateOutputOK(t, ctx, srv, ctrl, id, update)
	assert.Equal(t, id, out.ID, "")
	assert.Equal(t, "stdout", out.Name, "")
	assert.ContainsStr(t, "test", out.Tags, "")
	assert.ContainsStr(t, "foo", out.Tags, "")
	assert.Equal(t, true, out.Enabled, "")

	// DELETE
	test.DeleteOutputNoContent(t, ctx, srv, ctrl, id)

	// GET 404
	test.GetOutputNotFound(t, ctx, srv, ctrl, id)
}

func TestOutputDefs(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	ctrl := controller.NewOutputController(srv, db, om)
	ctx := context.Background()

	_, specs := test.SpecsOutputOK(t, ctx, srv, ctrl)
	assert.True(t, len(specs) > 0, "")
	for _, spec := range specs {
		if spec.Name == "http" {
			assert.True(t, len(spec.Props) == 1, "")
			assert.Equal(t, "url", spec.Props[0].Name, "")
			assert.Equal(t, "Target URL", spec.Props[0].Desc, "")
			assert.Equal(t, "string", spec.Props[0].Type, "")
		}
	}
}
