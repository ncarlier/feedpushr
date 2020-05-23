package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/autogen/app"
	"github.com/ncarlier/feedpushr/v3/autogen/app/test"
	"github.com/ncarlier/feedpushr/v3/pkg/controller"
)

func TestOutputCRUD(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	ctrl := controller.NewOutputController(srv, db, outputs)
	ctx := context.Background()

	// CREATE
	alias := "test output"
	payload := &app.CreateOutputPayload{
		Alias:     alias,
		Name:      "stdout",
		Condition: "\"test\" in Tags",
	}
	_, out := test.CreateOutputCreated(t, ctx, srv, ctrl, payload)
	assert.Equal(t, alias, out.Alias)
	assert.Equal(t, "stdout", out.Name)
	assert.False(t, out.Enabled)
	assert.Equal(t, "\"test\" in Tags", out.Condition)
	assert.Equal(t, 0, out.NbSuccess)
	id := out.ID

	// GET
	_, out = test.GetOutputOK(t, ctx, srv, ctrl, id)
	assert.Equal(t, id, out.ID)
	assert.Equal(t, "test output", out.Alias)
	assert.Equal(t, "stdout", out.Name)

	// FIND
	_, list := test.ListOutputOK(t, ctx, srv, ctrl)
	assert.NotEmpty(t, list)
	item := list[len(list)-1]
	assert.Equal(t, id, item.ID)

	// UPDATE
	alias = "updated output"
	update := &app.UpdateOutputPayload{
		Alias:   &alias,
		Enabled: true,
	}
	_, out = test.UpdateOutputOK(t, ctx, srv, ctrl, id, update)
	assert.Equal(t, id, out.ID)
	assert.Equal(t, alias, out.Alias)
	assert.Equal(t, "stdout", out.Name)
	assert.Equal(t, "\"test\" in Tags", out.Condition)
	assert.True(t, out.Enabled)

	// DELETE
	test.DeleteOutputNoContent(t, ctx, srv, ctrl, id)

	// GET 404
	test.GetOutputNotFound(t, ctx, srv, ctrl, id)
}

func TestOutputDefs(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	ctrl := controller.NewOutputController(srv, db, outputs)
	ctx := context.Background()

	_, specs := test.SpecsOutputOK(t, ctx, srv, ctrl)
	assert.NotEmpty(t, specs)
	for _, spec := range specs {
		if spec.Name == "http" {
			assert.Equal(t, 3, len(spec.Props))
			assert.Equal(t, "url", spec.Props[0].Name)
			assert.Equal(t, "Target URL", spec.Props[0].Desc)
			assert.Equal(t, "url", spec.Props[0].Type)
		}
	}
}

func TestFilterCRUD(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	ctrl := controller.NewOutputController(srv, db, outputs)
	ctx := context.Background()

	// CREATE OUTPUT
	alias := "test output"
	outputPayload := &app.CreateOutputPayload{
		Alias:     alias,
		Name:      "stdout",
		Condition: "\"test\" in Tags",
	}
	_, out := test.CreateOutputCreated(t, ctx, srv, ctrl, outputPayload)
	assert.Equal(t, alias, out.Alias)
	assert.Equal(t, "stdout", out.Name)
	assert.False(t, out.Enabled)
	assert.Equal(t, "\"test\" in Tags", out.Condition)
	assert.Equal(t, 0, out.NbSuccess)
	assert.Empty(t, out.Filters)
	id := out.ID

	// CREATE
	alias = "Add test prefix"
	payload := &app.CreateFilterOutputPayload{
		Alias: alias,
		Name:  "title",
		Props: map[string]interface{}{
			"prefix": "[test]",
		},
		Condition: "\"prefix\" in Tags",
	}
	_, f := test.CreateFilterOutputCreated(t, ctx, srv, ctrl, id, payload)
	assert.Equal(t, alias, f.Alias)
	assert.Equal(t, "title", f.Name)
	assert.False(t, f.Enabled)
	assert.Equal(t, "\"prefix\" in Tags", f.Condition)
	assert.Equal(t, "[test]", f.Props["prefix"])
	assert.NotEqual(t, "", f.ID)
	filterID := f.ID

	// Check output def
	_, out = test.GetOutputOK(t, ctx, srv, ctrl, id)
	assert.Equal(t, id, out.ID)
	assert.Equal(t, "stdout", out.Name)
	assert.Len(t, out.Filters, 1)
	assert.Equal(t, f.ID, out.Filters[0].ID)
	assert.Equal(t, f.Name, out.Filters[0].Name)
	assert.Equal(t, f.Condition, out.Filters[0].Condition)

	// UPDATE
	update := &app.UpdateFilterOutputPayload{
		Enabled: true,
	}
	_, f = test.UpdateFilterOutputOK(t, ctx, srv, ctrl, id, filterID, update)
	assert.Equal(t, "Add test prefix", f.Alias)
	assert.Equal(t, "title", f.Name)
	assert.Equal(t, "\"prefix\" in Tags", f.Condition)
	assert.True(t, f.Enabled)
	assert.Equal(t, filterID, f.ID)

	// DELETE
	test.DeleteFilterOutputNoContent(t, ctx, srv, ctrl, id, filterID)

	// TODO verify output def
}
