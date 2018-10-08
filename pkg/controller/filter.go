package controller

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/filter"
)

// FilterController implements the filter resource.
type FilterController struct {
	*goa.Controller
	cf *filter.Chain
}

// NewFilterController creates a filter controller.
func NewFilterController(service *goa.Service, cf *filter.Chain) *FilterController {
	return &FilterController{
		Controller: service.NewController("FilterController"),
		cf:         cf,
	}
}

// List runs the list action.
func (c *FilterController) List(ctx *app.ListFilterContext) error {
	res := app.FilterCollection{}
	specs := c.cf.GetSpec()
	for _, spec := range specs {
		f := app.Filter{
			Name:  spec.Name,
			Desc:  spec.Desc,
			Props: spec.Props,
			Tags:  spec.Tags,
		}
		res = append(res, &f)
	}
	return ctx.OK(res)
}
