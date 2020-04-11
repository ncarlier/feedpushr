package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
)

// ExploreController implements the explore resource.
type ExploreController struct {
	*goa.Controller
}

// NewExploreController creates a explore controller.
func NewExploreController(service *goa.Service) *ExploreController {
	return &ExploreController{Controller: service.NewController("ExploreController")}
}

// Get runs the get action.
func (c *ExploreController) Get(ctx *app.GetExploreContext) error {
	// ExploreController_Get: start_implement

	// Put your logic here

	res := app.ExploreResponseCollection{}
	return ctx.OK(res)
	// ExploreController_Get: end_implement
}
