package controller

import (
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/filter"
)

// FilterController implements the filter resource.
type FilterController struct {
	*goa.Controller
}

// NewFilterController creates a filter controller.
func NewFilterController(service *goa.Service) *FilterController {
	return &FilterController{
		Controller: service.NewController("FilterController"),
	}
}

// Specs runs the specs action.
func (c *FilterController) Specs(ctx *app.SpecsFilterContext) error {
	specs := filter.GetAvailableFilters()

	res := app.FilterSpecResponseCollection{}
	for _, spec := range specs {
		s := &app.FilterSpecResponse{
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
