package controller

import (
	"errors"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
	"github.com/ncarlier/feedpushr/v3/pkg/explore"
)

// ExploreController implements the explore resource.
type ExploreController struct {
	*goa.Controller
	explorer explore.Explorer
}

// NewExploreController creates a explore controller.
func NewExploreController(service *goa.Service, explorer explore.Explorer) *ExploreController {
	return &ExploreController{
		Controller: service.NewController("ExploreController"),
		explorer:   explorer,
	}
}

// Get runs the get action.
func (c *ExploreController) Get(ctx *app.GetExploreContext) error {
	if ctx.Q == nil || *ctx.Q == "" {
		err := errors.New("missing query parameter")
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	results, err := c.explorer.Search(*ctx.Q)
	if err != nil {
		return goa.ErrInternal(err)
	}
	res := app.ExploreResponseCollection{}
	for _, result := range *results {
		res = append(res, &app.ExploreResponse{
			Desc:    result.Desc,
			HTMLURL: result.HTMLURL,
			Title:   result.Title,
			XMLURL:  result.XMLURL,
		})

	}
	return ctx.OK(res)
}
