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

// Create runs the create action.
func (c *OutputController) Create(ctx *app.CreateOutputContext) error {
	// OutputController_Create: start_implement

	// Put your logic here

	return nil
	// OutputController_Create: end_implement
}

// Delete runs the delete action.
func (c *OutputController) Delete(ctx *app.DeleteOutputContext) error {
	// OutputController_Delete: start_implement

	// Put your logic here

	return nil
	// OutputController_Delete: end_implement
}

// Get runs the get action.
func (c *OutputController) Get(ctx *app.GetOutputContext) error {
	// OutputController_Get: start_implement

	// Put your logic here

	res := &app.Output{}
	return ctx.OK(res)
	// OutputController_Get: end_implement
}

// Update runs the update action.
func (c *OutputController) Update(ctx *app.UpdateOutputContext) error {
	// OutputController_Update: start_implement

	// Put your logic here

	res := &app.Output{}
	return ctx.OK(res)
	// OutputController_Update: end_implement
}

// List runs the list action.
func (c *OutputController) List(ctx *app.ListOutputContext) error {
	res := app.OutputCollection{}
	outputs := c.om.GetOutputDefs()
	for _, def := range outputs {
		o := app.Output{
			Name:  def.Name,
			Desc:  def.Desc,
			Props: def.Props,
			Tags:  def.Tags,
		}
		res = append(res, &o)
	}
	return ctx.OK(res)
}

// Specs runs the specs action.
func (c *OutputController) Specs(ctx *app.SpecsOutputContext) error {
	// OutputController_Specs: start_implement

	// Put your logic here

	res := app.OutputSpecCollection{}
	return ctx.OK(res)
	// OutputController_Specs: end_implement
}
