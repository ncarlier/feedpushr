package controller

import (
	"strings"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
	"github.com/ncarlier/feedpushr/v3/pkg/version"
)

// IndexController implements the index resource.
type IndexController struct {
	*goa.Controller
	issuer string
}

// NewIndexController creates a index controller.
func NewIndexController(service *goa.Service, issuer string) *IndexController {
	return &IndexController{
		Controller: service.NewController("IndexController"),
		issuer:     issuer,
	}
}

var links = make(map[string]*app.HALLink)

func init() {
	links["documentation"] = &app.HALLink{
		Href: "https://github.com/ncarlier/feedpushr",
	}
}

// Get runs the get action.
func (c *IndexController) Get(ctx *app.GetIndexContext) error {
	if _, ok := links["issuer"]; !ok && strings.HasPrefix(c.issuer, "http") {
		links["issuer"] = &app.HALLink{
			Href: c.issuer,
		}
	}
	res := &app.Info{
		Name:    "feedpushr",
		Desc:    "Feed aggregator daemon with sugar on top",
		Version: version.Version,
		Links:   links,
	}
	return ctx.OK(res)
}
