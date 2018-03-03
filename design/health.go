package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("health", func() {

	Action("get", func() {
		Routing(
			GET("/healthz"),
		)
		Description("Perform health check.")
		Response(OK, "text/plain")
	})
})
