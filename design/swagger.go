package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Action("get", func() {
		Routing(
			GET("/swagger.json"),
		)
		Description("Get OpenAPI specifications")
		Response(OK, func() {
			Media("application/json")
		})
	})
})
