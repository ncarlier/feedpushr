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

// List runs the list action.
func (c *OutputController) List(ctx *app.ListOutputContext) error {
	res := app.OutputCollection{}
	specs := c.om.GetSpec()
	for _, spec := range specs {
		o := app.Output{
			Name:  spec.Name,
			Desc:  spec.Desc,
			Props: spec.Props,
			Tags:  spec.Tags,
		}
		res = append(res, &o)
	}
	return ctx.OK(res)
}
