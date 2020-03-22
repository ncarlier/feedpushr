package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("opml", func() {

	BasePath("/opml")

	Action("get", func() {
		Routing(
			GET(""),
		)
		Description("Get all feeds as an OMPL format")
		Response(OK, func() {
			Media("application/xml")
		})
		Response(BadRequest, ErrorMedia)
	})

	Action("upload", func() {
		Routing(
			POST(""),
		)
		Description("Upload an OPML file to create feeds")
		Response(Created, "/feeds")
		Response(BadRequest, ErrorMedia)
	})
})
