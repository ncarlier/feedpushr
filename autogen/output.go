package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
)

// OutputController implements the output resource.
type OutputController struct {
	*goa.Controller
}

// NewOutputController creates a output controller.
func NewOutputController(service *goa.Service) *OutputController {
	return &OutputController{Controller: service.NewController("OutputController")}
}

// Create runs the create action.
func (c *OutputController) Create(ctx *app.CreateOutputContext) error {
	// OutputController_Create: start_implement

	// Put your logic here

	return nil
	// OutputController_Create: end_implement
}

// Delete runs the delete action.
func (c *OutputController) Delete(ctx *app.DeleteOutputContext) error {
	// OutputController_Delete: start_implement

	// Put your logic here

	return nil
	// OutputController_Delete: end_implement
}

// Get runs the get action.
func (c *OutputController) Get(ctx *app.GetOutputContext) error {
	// OutputController_Get: start_implement

	// Put your logic here

	res := &app.Output{}
	return ctx.OK(res)
	// OutputController_Get: end_implement
}

// List runs the list action.
func (c *OutputController) List(ctx *app.ListOutputContext) error {
	// OutputController_List: start_implement

	// Put your logic here

	res := app.OutputCollection{}
	return ctx.OK(res)
	// OutputController_List: end_implement
}

// Specs runs the specs action.
func (c *OutputController) Specs(ctx *app.SpecsOutputContext) error {
	// OutputController_Specs: start_implement

	// Put your logic here

	res := app.OutputSpecCollection{}
	return ctx.OK(res)
	// OutputController_Specs: end_implement
}

// Update runs the update action.
func (c *OutputController) Update(ctx *app.UpdateOutputContext) error {
	// OutputController_Update: start_implement

	// Put your logic here

	res := &app.Output{}
	return ctx.OK(res)
	// OutputController_Update: end_implement
}
