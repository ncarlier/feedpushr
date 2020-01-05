package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/autogen/app"
)

// IndexController implements the index resource.
type IndexController struct {
	*goa.Controller
}

// NewIndexController creates a index controller.
func NewIndexController(service *goa.Service) *IndexController {
	return &IndexController{Controller: service.NewController("IndexController")}
}

// Get runs the get action.
func (c *IndexController) Get(ctx *app.GetIndexContext) error {
	// IndexController_Get: start_implement

	// Put your logic here

	res := &app.Info{}
	return ctx.OK(res)
	// IndexController_Get: end_implement
}
