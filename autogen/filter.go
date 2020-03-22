package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
)

// FilterController implements the filter resource.
type FilterController struct {
	*goa.Controller
}

// NewFilterController creates a filter controller.
func NewFilterController(service *goa.Service) *FilterController {
	return &FilterController{Controller: service.NewController("FilterController")}
}

// Create runs the create action.
func (c *FilterController) Create(ctx *app.CreateFilterContext) error {
	// FilterController_Create: start_implement

	// Put your logic here

	return nil
	// FilterController_Create: end_implement
}

// Delete runs the delete action.
func (c *FilterController) Delete(ctx *app.DeleteFilterContext) error {
	// FilterController_Delete: start_implement

	// Put your logic here

	return nil
	// FilterController_Delete: end_implement
}

// Get runs the get action.
func (c *FilterController) Get(ctx *app.GetFilterContext) error {
	// FilterController_Get: start_implement

	// Put your logic here

	res := &app.FilterResponse{}
	return ctx.OK(res)
	// FilterController_Get: end_implement
}

// List runs the list action.
func (c *FilterController) List(ctx *app.ListFilterContext) error {
	// FilterController_List: start_implement

	// Put your logic here

	res := app.FilterResponseCollection{}
	return ctx.OK(res)
	// FilterController_List: end_implement
}

// Specs runs the specs action.
func (c *FilterController) Specs(ctx *app.SpecsFilterContext) error {
	// FilterController_Specs: start_implement

	// Put your logic here

	res := app.FilterSpecResponseCollection{}
	return ctx.OK(res)
	// FilterController_Specs: end_implement
}

// Update runs the update action.
func (c *FilterController) Update(ctx *app.UpdateFilterContext) error {
	// FilterController_Update: start_implement

	// Put your logic here

	res := &app.FilterResponse{}
	return ctx.OK(res)
	// FilterController_Update: end_implement
}
