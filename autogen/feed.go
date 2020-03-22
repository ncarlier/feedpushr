package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
)

// FeedController implements the feed resource.
type FeedController struct {
	*goa.Controller
}

// NewFeedController creates a feed controller.
func NewFeedController(service *goa.Service) *FeedController {
	return &FeedController{Controller: service.NewController("FeedController")}
}

// Create runs the create action.
func (c *FeedController) Create(ctx *app.CreateFeedContext) error {
	// FeedController_Create: start_implement

	// Put your logic here

	return nil
	// FeedController_Create: end_implement
}

// Delete runs the delete action.
func (c *FeedController) Delete(ctx *app.DeleteFeedContext) error {
	// FeedController_Delete: start_implement

	// Put your logic here

	return nil
	// FeedController_Delete: end_implement
}

// Get runs the get action.
func (c *FeedController) Get(ctx *app.GetFeedContext) error {
	// FeedController_Get: start_implement

	// Put your logic here

	res := &app.FeedResponse{}
	return ctx.OK(res)
	// FeedController_Get: end_implement
}

// List runs the list action.
func (c *FeedController) List(ctx *app.ListFeedContext) error {
	// FeedController_List: start_implement

	// Put your logic here

	res := &app.FeedsPageResponse{}
	return ctx.OK(res)
	// FeedController_List: end_implement
}

// Start runs the start action.
func (c *FeedController) Start(ctx *app.StartFeedContext) error {
	// FeedController_Start: start_implement

	// Put your logic here

	return nil
	// FeedController_Start: end_implement
}

// Stop runs the stop action.
func (c *FeedController) Stop(ctx *app.StopFeedContext) error {
	// FeedController_Stop: start_implement

	// Put your logic here

	return nil
	// FeedController_Stop: end_implement
}

// Update runs the update action.
func (c *FeedController) Update(ctx *app.UpdateFeedContext) error {
	// FeedController_Update: start_implement

	// Put your logic here

	res := &app.FeedResponse{}
	return ctx.OK(res)
	// FeedController_Update: end_implement
}
