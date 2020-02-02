package main

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
)

// OpmlController implements the opml resource.
type OpmlController struct {
	*goa.Controller
}

// NewOpmlController creates a opml controller.
func NewOpmlController(service *goa.Service) *OpmlController {
	return &OpmlController{Controller: service.NewController("OpmlController")}
}

// Get runs the get action.
func (c *OpmlController) Get(ctx *app.GetOpmlContext) error {
	// OpmlController_Get: start_implement

	// Put your logic here

	return nil
	// OpmlController_Get: end_implement
}

// Upload runs the upload action.
func (c *OpmlController) Upload(ctx *app.UploadOpmlContext) error {
	// OpmlController_Upload: start_implement

	// Put your logic here

	return nil
	// OpmlController_Upload: end_implement
}
