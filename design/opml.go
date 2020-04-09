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
		Description("Get all feeds as an OPML format")
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
		Response(Accepted, OPMLImportJobResponse)
		Response(BadRequest, ErrorMedia)
	})

	Action("status", func() {
		Routing(
			GET("/status/:id"),
		)
		Description("Get OPML import status")
		Params(func() {
			Param("id", Integer, "Import job ID")
		})
		Response(OK, func() {
			Media("text/event-stream")
		})
		Response(NotFound, ErrorMedia)
	})
})

// OPMLImportJobResponse is OPM import job media type.
var OPMLImportJobResponse = MediaType("application/vnd.feedpushr.ompl-import-job.v2+json", func() {
	TypeName("OPMLImportJobResponse")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", String, "ID of the import job")
		Required("id")
	})

	View("default", func() {
		Attribute("id")
	})
})
