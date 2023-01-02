package controller

import (
	"github.com/ncarlier/feedpushr/v3/autogen/app"
	"github.com/ncarlier/feedpushr/v3/pkg/filter"
)

// CreateFilter runs the createFilter action.
func (c *OutputController) CreateFilter(ctx *app.CreateFilterOutputContext) error {
	processor, err := c.outputs.GetOutputProcessor(ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	filterDef := filter.NewBuilder().Alias(
		&ctx.Payload.Alias,
	).Spec(
		ctx.Payload.Name,
	).Props(
		ctx.Payload.Props,
	).Condition(
		&ctx.Payload.Condition,
	).Enable(false).NewID().Build()

	f, err := processor.Filters.Add(filterDef)
	if err != nil {
		return err
	}
	_, err = c.db.SaveOutput(processor.GetDef())
	if err != nil {
		return err
	}
	def := f.GetDef()
	return ctx.Created(filter.NewFilterResponseFromDef(&def))
}

// DeleteFilter runs the deleteFilter action.
func (c *OutputController) DeleteFilter(ctx *app.DeleteFilterOutputContext) error {
	processor, err := c.outputs.GetOutputProcessor(ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}
	if err := processor.Filters.Remove(ctx.IDFilter); err != nil {
		return err
	}
	_, err = c.db.SaveOutput(processor.GetDef())
	if err != nil {
		return err
	}
	return ctx.NoContent()
}

// UpdateFilter runs the updateFilter action.
func (c *OutputController) UpdateFilter(ctx *app.UpdateFilterOutputContext) error {
	processor, err := c.outputs.GetOutputProcessor(ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}
	f, err := processor.Filters.Get(ctx.IDFilter)
	if err != nil {
		return ctx.NotFound()
	}
	filterDef := filter.NewBuilder().From(
		f.GetDef(),
	).Alias(
		ctx.Payload.Alias,
	).Props(
		ctx.Payload.Props,
	).Condition(
		ctx.Payload.Condition,
	).Enable(
		ctx.Payload.Enabled,
	).ID(ctx.IDFilter).Build()
	f, err = processor.Filters.Update(ctx.IDFilter, filterDef)
	if err != nil {
		return err
	}
	_, err = c.db.SaveOutput(processor.GetDef())
	if err != nil {
		return err
	}
	def := f.GetDef()
	return ctx.OK(filter.NewFilterResponseFromDef(&def))
}
