package controller

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/builder"
	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/filter"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
)

// FilterController implements the filter resource.
type FilterController struct {
	*goa.Controller
	cf *filter.Chain
	db store.DB
}

// NewFilterController creates a filter controller.
func NewFilterController(service *goa.Service, db store.DB, cf *filter.Chain) *FilterController {
	return &FilterController{
		Controller: service.NewController("FilterController"),
		cf:         cf,
		db:         db,
	}
}

// Create runs the create action.
func (c *FilterController) Create(ctx *app.CreateFilterContext) error {
	filter := builder.NewFilterBuilder().Alias(
		&ctx.Payload.Alias,
	).Spec(
		ctx.Payload.Name,
	).Props(
		ctx.Payload.Props,
	).Condition(
		&ctx.Payload.Condition,
	).Enable(false).Build()

	f, err := c.cf.Add(filter)
	if err != nil {
		return err
	}

	def, err := c.db.SaveFilter(f.GetDef())
	if err != nil {
		return err
	}

	return ctx.Created(builder.NewFilterResponseFromDef(def))
}

// Delete runs the delete action.
func (c *FilterController) Delete(ctx *app.DeleteFilterContext) error {
	filter := &model.FilterDef{
		ID: ctx.ID,
	}
	err := c.cf.Remove(filter)
	if err != nil {
		return ctx.NotFound()
	}

	_, err = c.db.DeleteFilter(filter.ID)
	if err != nil {
		return err
	}

	return ctx.NoContent()
}

// Get runs the get action.
func (c *FilterController) Get(ctx *app.GetFilterContext) error {
	filter, err := c.cf.Get(ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	def := filter.GetDef()
	return ctx.OK(builder.NewFilterResponseFromDef(&def))
}

// Update runs the update action.
func (c *FilterController) Update(ctx *app.UpdateFilterContext) error {
	f, err := c.cf.Get(ctx.ID)
	if err != nil {
		if err == common.ErrFilterNotFound {
			return ctx.NotFound()
		}
		return err
	}

	update := builder.NewFilterBuilder().From(
		f.GetDef(),
	).Alias(
		ctx.Payload.Alias,
	).Props(
		ctx.Payload.Props,
	).Condition(
		ctx.Payload.Condition,
	).Enable(
		ctx.Payload.Enabled,
	).Build()

	f, err = c.cf.Update(update)
	if err != nil {
		return err
	}

	def, err := c.db.SaveFilter(f.GetDef())
	if err != nil {
		return err
	}

	return ctx.OK(builder.NewFilterResponseFromDef(def))
}

// List runs the list action.
func (c *FilterController) List(ctx *app.ListFilterContext) error {
	res := app.FilterResponseCollection{}
	filters := c.cf.GetFilterDefs()
	for _, def := range filters {
		res = append(res, builder.NewFilterResponseFromDef(&def))
	}
	return ctx.OK(res)
}

// Specs runs the specs action.
func (c *FilterController) Specs(ctx *app.SpecsFilterContext) error {
	specs := filter.GetAvailableFilters()

	res := app.FilterSpecResponseCollection{}
	for _, spec := range specs {
		s := &app.FilterSpecResponse{
			Name:  spec.Name,
			Desc:  spec.Desc,
			Props: app.PropSpecCollection{},
		}
		for _, prop := range spec.PropsSpec {
			s.Props = append(s.Props, &app.PropSpec{
				Name: prop.Name,
				Desc: prop.Desc,
				Type: prop.Type.String(),
			})
		}

		res = append(res, s)
	}

	return ctx.OK(res)
}
