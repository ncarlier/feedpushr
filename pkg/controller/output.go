package controller

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
	"github.com/ncarlier/feedpushr/v3/pkg/builder"
	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/output"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
)

// OutputController implements the output resource.
type OutputController struct {
	*goa.Controller
	outputs *output.Manager
	db      store.DB
}

// NewOutputController creates a output controller.
func NewOutputController(service *goa.Service, db store.DB, outputs *output.Manager) *OutputController {
	return &OutputController{
		Controller: service.NewController("OutputController"),
		outputs:    outputs,
		db:         db,
	}
}

// Create runs the create action.
func (c *OutputController) Create(ctx *app.CreateOutputContext) error {
	def := builder.NewOutputBuilder().Alias(
		&ctx.Payload.Alias,
	).Spec(
		ctx.Payload.Name,
	).Props(
		ctx.Payload.Props,
	).Condition(
		&ctx.Payload.Condition,
	).Enable(false).NewID().Build()
	processor, err := c.outputs.AddOutputProcessor(def)
	if err != nil {
		return err
	}
	def, err = c.db.SaveOutput(processor.GetDef())
	if err != nil {
		// cleanup previous created processor
		_def := processor.GetDef()
		c.outputs.RemoveOutputProcessor(&_def)
		return err
	}
	return ctx.Created(builder.NewOutputResponseFromDef(def))
}

// Delete runs the delete action.
func (c *OutputController) Delete(ctx *app.DeleteOutputContext) error {
	def := &model.OutputDef{
		ID: ctx.ID,
	}
	err := c.outputs.RemoveOutputProcessor(def)
	if err != nil {
		return ctx.NotFound()
	}
	_, err = c.db.DeleteOutput(def.ID)
	if err != nil {
		return err
	}

	return ctx.NoContent()
}

// Get runs the get action.
func (c *OutputController) Get(ctx *app.GetOutputContext) error {
	processor, err := c.outputs.GetOutputProcessor(ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	def := processor.GetDef()
	return ctx.OK(builder.NewOutputResponseFromDef(&def))
}

// Update runs the update action.
func (c *OutputController) Update(ctx *app.UpdateOutputContext) error {
	processor, err := c.outputs.GetOutputProcessor(ctx.ID)
	if err != nil {
		if err == common.ErrOutputNotFound {
			return ctx.NotFound()
		}
		return err
	}

	update := builder.NewOutputBuilder().From(
		processor.GetDef(),
	).Alias(
		ctx.Payload.Alias,
	).Props(
		ctx.Payload.Props,
	).Condition(
		ctx.Payload.Condition,
	).Enable(
		ctx.Payload.Enabled,
	).Build()

	processor, err = c.outputs.UpdateOutputProcessor(update)
	if err != nil {
		return err
	}

	def, err := c.db.SaveOutput(processor.GetDef())
	if err != nil {
		return err
	}

	return ctx.OK(builder.NewOutputResponseFromDef(def))
}

// List runs the list action.
func (c *OutputController) List(ctx *app.ListOutputContext) error {
	res := app.OutputResponseCollection{}
	outputs := c.outputs.GetOutputDefs()
	for _, def := range outputs {
		res = append(res, builder.NewOutputResponseFromDef(&def))
	}
	return ctx.OK(res)
}

// Specs runs the specs action.
func (c *OutputController) Specs(ctx *app.SpecsOutputContext) error {
	specs := c.outputs.GetAvailableOutputs()

	res := app.OutputSpecResponseCollection{}
	for _, spec := range specs {
		s := &app.OutputSpecResponse{
			Name:  spec.Name,
			Desc:  spec.Desc,
			Props: app.PropSpecCollection{},
		}
		for _, prop := range spec.PropsSpec {
			s.Props = append(s.Props, &app.PropSpec{
				Name:    prop.Name,
				Desc:    prop.Desc,
				Type:    prop.Type.String(),
				Options: prop.Options,
			})
		}

		res = append(res, s)
	}

	return ctx.OK(res)
}
