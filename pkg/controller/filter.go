package controller

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
	"github.com/ncarlier/feedpushr/v3/pkg/filter"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// FilterController implements the filter resource.
type FilterController struct {
	*goa.Controller
	specs []model.Spec
}

// NewFilterController creates a filter controller.
func NewFilterController(service *goa.Service, chain *filter.Chain) *FilterController {
	return &FilterController{
		Controller: service.NewController("FilterController"),
		specs:      chain.GetAvailableFilters(),
	}
}

// Specs runs the specs action.
func (c *FilterController) Specs(ctx *app.SpecsFilterContext) error {
	res := app.FilterSpecResponseCollection{}
	for _, spec := range c.specs {
		s := &app.FilterSpecResponse{
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
