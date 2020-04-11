package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
)

// VarsController implements the vars resource.
type VarsController struct {
	*goa.Controller
}

// NewVarsController creates a vars controller.
func NewVarsController(service *goa.Service) *VarsController {
	return &VarsController{Controller: service.NewController("VarsController")}
}

// Get runs the get action.
func (c *VarsController) Get(ctx *app.GetVarsContext) error {
	// VarsController_Get: start_implement

	// Put your logic here

	return nil
	// VarsController_Get: end_implement
}
