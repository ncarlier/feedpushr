package controller

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/filter"
)

// FilterController implements the filter resource.
type FilterController struct {
	*goa.Controller
	cf *filter.Chain
}

// NewFilterController creates a filter controller.
func NewFilterController(service *goa.Service, cf *filter.Chain) *FilterController {
	return &FilterController{
		Controller: service.NewController("FilterController"),
		cf:         cf,
	}
}

// Create runs the create action.
func (c *FilterController) Create(ctx *app.CreateFilterContext) error {
	filter := &app.Filter{
		Name:    ctx.Payload.Name,
		Props:   ctx.Payload.Props,
		Tags:    builder.GetFeedTags(ctx.Payload.Tags),
		Enabled: false,
	}
	f, err := c.cf.Add(filter)
	if err != nil {
		return err
	}
	res := builder.NewFilterFromDef(f.GetDef())
	return ctx.Created(res)
}

// Delete runs the delete action.
func (c *FilterController) Delete(ctx *app.DeleteFilterContext) error {
	filter := &app.Filter{
		ID: ctx.ID,
	}
	err := c.cf.Remove(filter)
	if err != nil {
		return ctx.NotFound()
	}
	return ctx.NoContent()
}

// Get runs the get action.
func (c *FilterController) Get(ctx *app.GetFilterContext) error {
	filter, err := c.cf.Get(ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	res := builder.NewFilterFromDef(filter.GetDef())
	return ctx.OK(res)
}

// Update runs the update action.
func (c *FilterController) Update(ctx *app.UpdateFilterContext) error {
	update := &app.Filter{
		ID:      ctx.ID,
		Props:   ctx.Payload.Props,
		Tags:    builder.GetFeedTags(ctx.Payload.Tags),
		Enabled: ctx.Payload.Enabled,
	}
	f, err := c.cf.Update(update)
	if err != nil {
		return err
	}

	res := builder.NewFilterFromDef(f.GetDef())
	return ctx.OK(res)
}

// List runs the list action.
func (c *FilterController) List(ctx *app.ListFilterContext) error {
	res := app.FilterCollection{}
	filters := c.cf.GetFilterDefs()
	for _, def := range filters {
		res = append(res, builder.NewFilterFromDef(def))
	}
	return ctx.OK(res)
}

// Specs runs the specs action.
func (c *FilterController) Specs(ctx *app.SpecsFilterContext) error {
	specs := filter.GetAvailableFilters()

	res := app.FilterSpecCollection{}
	for _, spec := range specs {
		s := &app.FilterSpec{
			Name:  spec.Name,
			Desc:  spec.Desc,
			Props: app.PropSpecCollection{},
		}
		for _, prop := range spec.PropsSpec {
			s.Props = append(s.Props, &app.PropSpec{
				Name: prop.Name,
				Desc: prop.Desc,
				Type: prop.Type,
			})
		}

		res = append(res, s)
	}

	return ctx.OK(res)
}
