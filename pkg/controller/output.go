package controller

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/feedpushr/pkg/output"
	"github.com/ncarlier/feedpushr/pkg/store"
)

// OutputController implements the output resource.
type OutputController struct {
	*goa.Controller
	om *output.Manager
	db store.DB
}

// NewOutputController creates a output controller.
func NewOutputController(service *goa.Service, db store.DB, om *output.Manager) *OutputController {
	return &OutputController{
		Controller: service.NewController("OutputController"),
		om:         om,
		db:         db,
	}
}

// Create runs the create action.
func (c *OutputController) Create(ctx *app.CreateOutputContext) error {
	out := builder.NewOutputBuilder().Spec(
		ctx.Payload.Name,
	).Props(
		ctx.Payload.Props,
	).Tags(
		ctx.Payload.Tags,
	).Enable(false).Build()
	provider, err := c.om.Add(out)
	if err != nil {
		return err
	}
	def, err := c.db.SaveOutput(provider.GetDef())
	if err != nil {
		return err
	}
	res := builder.NewOutputFromDef(*def)
	return ctx.Created(res)
}

// Delete runs the delete action.
func (c *OutputController) Delete(ctx *app.DeleteOutputContext) error {
	out := &model.OutputDef{
		ID: ctx.ID,
	}
	err := c.om.Remove(out)
	if err != nil {
		return ctx.NotFound()
	}
	_, err = c.db.DeleteOutput(out.ID)
	if err != nil {
		return err
	}

	return ctx.NoContent()
}

// Get runs the get action.
func (c *OutputController) Get(ctx *app.GetOutputContext) error {
	out, err := c.om.Get(ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	res := builder.NewOutputFromDef(out.GetDef())
	return ctx.OK(res)
}

// Update runs the update action.
func (c *OutputController) Update(ctx *app.UpdateOutputContext) error {
	update := builder.NewOutputBuilder().ID(
		ctx.ID,
	).Props(
		ctx.Payload.Props,
	).Tags(
		ctx.Payload.Tags,
	).Enable(
		ctx.Payload.Enabled,
	).Build()
	out, err := c.om.Update(update)
	if err != nil {
		return err
	}

	def, err := c.db.SaveOutput(out.GetDef())
	if err != nil {
		return err
	}

	res := builder.NewOutputFromDef(*def)
	return ctx.OK(res)
}

// List runs the list action.
func (c *OutputController) List(ctx *app.ListOutputContext) error {
	res := app.OutputCollection{}
	outputs := c.om.GetOutputDefs()
	for _, def := range outputs {
		res = append(res, builder.NewOutputFromDef(def))
	}
	return ctx.OK(res)
}

// Specs runs the specs action.
func (c *OutputController) Specs(ctx *app.SpecsOutputContext) error {
	specs := output.GetAvailableOutputs()

	res := app.OutputSpecCollection{}
	for _, spec := range specs {
		s := &app.OutputSpec{
			Name:  spec.Name,
			Desc:  spec.Desc,
			Props: app.PropSpecCollection{},
		}
		for _, prop := range spec.PropsSpec {
			s.Props = append(s.Props, &app.PropSpec{
				Name: prop.Name,
				Desc: prop.Desc,
				Type: prop.Type.String(),
			})
		}

		res = append(res, s)
	}

	return ctx.OK(res)
}
