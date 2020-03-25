package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("filter", func() {
	DefaultMedia(FilterResponse)
	BasePath("/filters")

	Action("specs", func() {
		Routing(
			GET("/_specs"),
		)
		Description("Retrieve all filter types available")
		Response(OK, func() {
			Media(CollectionOf(FilterSpecResponse, func() {
				View("default")
			}))
		})
	})
})

// FilterSpecResponse is the filter specification media type.
var FilterSpecResponse = MediaType("application/vnd.feedpushr.filter-spec.v2+json", func() {
	Description("The filter specification")
	TypeName("FilterSpecResponse")
	ContentType("application/json")
	Attributes(func() {
		Attribute("name", String, "Name of the filter", func() {
			Example("title")
		})
		Attribute("desc", String, "Description of the filter", func() {
			Example("Add a prefix to the tittle...")
		})
		Attribute("props", CollectionOf("application/vnd.feedpushr.prop-spec.v2+json"))
		Required("name", "desc", "props")
	})

	View("default", func() {
		Attribute("name")
		Attribute("desc")
		Attribute("props")
	})
})
