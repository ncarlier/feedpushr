package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
)

// SwaggerController implements the swagger resource.
type SwaggerController struct {
	*goa.Controller
}

// NewSwaggerController creates a swagger controller.
func NewSwaggerController(service *goa.Service) *SwaggerController {
	return &SwaggerController{Controller: service.NewController("SwaggerController")}
}

// Get runs the get action.
func (c *SwaggerController) Get(ctx *app.GetSwaggerContext) error {
	// SwaggerController_Get: start_implement

	// Put your logic here

	return nil
	// SwaggerController_Get: end_implement
}
