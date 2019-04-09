package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/autogen/app"
)

// OutputController implements the output resource.
type OutputController struct {
	*goa.Controller
}

// NewOutputController creates a output controller.
func NewOutputController(service *goa.Service) *OutputController {
	return &OutputController{Controller: service.NewController("OutputController")}
}

// List runs the list action.
func (c *OutputController) List(ctx *app.ListOutputContext) error {
	// OutputController_List: start_implement

	// Put your logic here

	res := app.OutputCollection{}
	return ctx.OK(res)
	// OutputController_List: end_implement
}
