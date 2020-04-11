package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service) *HealthController {
	return &HealthController{Controller: service.NewController("HealthController")}
}

// Get runs the get action.
func (c *HealthController) Get(ctx *app.GetHealthContext) error {
	// HealthController_Get: start_implement

	// Put your logic here

	return nil
	// HealthController_Get: end_implement
}
