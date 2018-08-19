package controller

import (
	"io/ioutil"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/assets"
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
	file, err := assets.GetFS().Open("/swagger.json")
	if err != nil {
		return goa.ErrInternal(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return goa.ErrInternal(err)
	}
	return ctx.OK(bytes)
}
