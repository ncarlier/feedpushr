package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("vars", func() {
	BasePath("/vars")

	Action("get", func() {
		Routing(
			GET(""),
		)
		Description("Get all internals exp vars")
		Response(OK, func() {
			Media("application/json")
		})
	})
})
