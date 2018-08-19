package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/autogen/app"
)

// FilterController implements the filter resource.
type FilterController struct {
	*goa.Controller
}

// NewFilterController creates a filter controller.
func NewFilterController(service *goa.Service) *FilterController {
	return &FilterController{Controller: service.NewController("FilterController")}
}

// List runs the list action.
func (c *FilterController) List(ctx *app.ListFilterContext) error {
	// FilterController_List: start_implement

	// Put your logic here

	res := app.FilterCollection{}
	return ctx.OK(res)
	// FilterController_List: end_implement
}
