package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
)

// PshbController implements the pshb resource.
type PshbController struct {
	*goa.Controller
}

// NewPshbController creates a pshb controller.
func NewPshbController(service *goa.Service) *PshbController {
	return &PshbController{Controller: service.NewController("PshbController")}
}

// Pub runs the pub action.
func (c *PshbController) Pub(ctx *app.PubPshbContext) error {
	// PshbController_Pub: start_implement

	// Put your logic here

	return nil
	// PshbController_Pub: end_implement
}

// Sub runs the sub action.
func (c *PshbController) Sub(ctx *app.SubPshbContext) error {
	// PshbController_Sub: start_implement

	// Put your logic here

	return nil
	// PshbController_Sub: end_implement
}
