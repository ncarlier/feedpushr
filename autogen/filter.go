package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
)

// FilterController implements the filter resource.
type FilterController struct {
	*goa.Controller
}

// NewFilterController creates a filter controller.
func NewFilterController(service *goa.Service) *FilterController {
	return &FilterController{Controller: service.NewController("FilterController")}
}

// Specs runs the specs action.
func (c *FilterController) Specs(ctx *app.SpecsFilterContext) error {
	// FilterController_Specs: start_implement

	// Put your logic here

	res := app.FilterSpecResponseCollection{}
	return ctx.OK(res)
	// FilterController_Specs: end_implement
}
