package controller

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/builder"
	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
	"github.com/ncarlier/feedpushr/v2/pkg/pipeline"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
)

// OutputController implements the output resource.
type OutputController struct {
	*goa.Controller
	pipeline *pipeline.Pipeline
	db       store.DB
}

// NewOutputController creates a output controller.
func NewOutputController(service *goa.Service, db store.DB, pipe *pipeline.Pipeline) *OutputController {
	return &OutputController{
		Controller: service.NewController("OutputController"),
		pipeline:   pipe,
		db:         db,
	}
}

// Create runs the create action.
func (c *OutputController) Create(ctx *app.CreateOutputContext) error {
	out := builder.NewOutputBuilder().Alias(
		&ctx.Payload.Alias,
	).Spec(
		ctx.Payload.Name,
	).Props(
		ctx.Payload.Props,
	).Condition(
		&ctx.Payload.Condition,
	).Enable(false).Build()
	provider, err := c.pipeline.AddOutput(out)
	if err != nil {
		return err
	}
	def, err := c.db.SaveOutput(provider.GetDef())
	if err != nil {
		return err
	}
	return ctx.Created(builder.NewOutputResponseFromDef(def))
}

// Delete runs the delete action.
func (c *OutputController) Delete(ctx *app.DeleteOutputContext) error {
	out := &model.OutputDef{
		ID: ctx.ID,
	}
	err := c.pipeline.RemoveOutput(out)
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
	out, err := c.pipeline.GetOutput(ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	def := out.GetDef()
	return ctx.OK(builder.NewOutputResponseFromDef(&def))
}

// Update runs the update action.
func (c *OutputController) Update(ctx *app.UpdateOutputContext) error {
	out, err := c.pipeline.GetOutput(ctx.ID)
	if err != nil {
		if err == common.ErrOutputNotFound {
			return ctx.NotFound()
		}
		return err
	}

	update := builder.NewOutputBuilder().From(
		out.GetDef(),
	).Alias(
		ctx.Payload.Alias,
	).Props(
		ctx.Payload.Props,
	).Condition(
		ctx.Payload.Condition,
	).Enable(
		ctx.Payload.Enabled,
	).Build()

	out, err = c.pipeline.UpdateOutput(update)
	if err != nil {
		return err
	}

	def, err := c.db.SaveOutput(out.GetDef())
	if err != nil {
		return err
	}

	return ctx.OK(builder.NewOutputResponseFromDef(def))
}

// List runs the list action.
func (c *OutputController) List(ctx *app.ListOutputContext) error {
	res := app.OutputResponseCollection{}
	outputs := c.pipeline.GetOutputDefs()
	for _, def := range outputs {
		res = append(res, builder.NewOutputResponseFromDef(&def))
	}
	return ctx.OK(res)
}

// Specs runs the specs action.
func (c *OutputController) Specs(ctx *app.SpecsOutputContext) error {
	specs := c.pipeline.GetAvailableOutputs()

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
