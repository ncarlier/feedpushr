package controller

import (
	"expvar"
	"fmt"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
)

// VarsController implements the vars resource.
type VarsController struct {
	*goa.Controller
}

// NewVarsController creates a vars controller.
func NewVarsController(service *goa.Service) *VarsController {
	return &VarsController{Controller: service.NewController("VarsController")}
}

// Get returns all exp vars.
func (c *VarsController) Get(ctx *app.GetVarsContext) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json; charset=utf-8")
	w := ctx.ResponseWriter
	ctx.ResponseData.WriteHeader(200)
	fmt.Fprintf(w, "{\n")
	first := true
	expvar.Do(func(kv expvar.KeyValue) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
	return nil
}
