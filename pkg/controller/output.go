package controller

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/output"
)

// OutputController implements the output resource.
type OutputController struct {
	*goa.Controller
	om *output.Manager
}

// NewOutputController creates a output controller.
func NewOutputController(service *goa.Service, om *output.Manager) *OutputController {
	return &OutputController{
		Controller: service.NewController("OutputController"),
		om:         om,
	}
}

// Get runs the get action.
func (c *OutputController) Get(ctx *app.GetOutputContext) error {
	spec := c.om.GetSpec()
	res := &app.Output{
		Name:  spec.Name,
		Desc:  spec.Desc,
		Props: spec.Props,
	}

	return ctx.OK(res)
}
