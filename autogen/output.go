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

// Get runs the get action.
func (c *OutputController) Get(ctx *app.GetOutputContext) error {
	// OutputController_Get: start_implement

	// Put your logic here

	res := &app.Output{}
	return ctx.OK(res)
	// OutputController_Get: end_implement
}
