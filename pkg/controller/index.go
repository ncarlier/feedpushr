package controller

import (
	"expvar"
	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/version"
)

// IndexController implements the index resource.
type IndexController struct {
	*goa.Controller
}

// NewIndexController creates a index controller.
func NewIndexController(service *goa.Service) *IndexController {
	return &IndexController{Controller: service.NewController("IndexController")}
}

var links = make(map[string]*app.HALLink)

func init() {
	links["documentation"] = &app.HALLink{
		Href: "https://github.com/ncarlier/feedpushr",
	}
}

// Get runs the get action.
func (c *IndexController) Get(ctx *app.GetIndexContext) error {
	expvar.Get("version").String()
	res := &app.Info{
		Name:    "feedpushr",
		Desc:    "Feed aggregator daemon with sugar on top",
		Version: version.Version,
		Links:   links,
	}
	return ctx.OK(res)
}
