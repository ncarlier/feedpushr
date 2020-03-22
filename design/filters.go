package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("filter", func() {
	DefaultMedia(FilterResponse)
	BasePath("/filters")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all filters definitions")
		Response(OK, func() {
			Media(CollectionOf(FilterResponse, func() {
				View("default")
			}))
		})
	})

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

	Action("get", func() {
		Routing(
			GET("/:id"),
		)
		Description("Retrieve filter with given ID")
		Params(func() {
			Param("id", Integer, "Filter ID")
		})
		Response(OK, FilterResponse)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create a new filter")
		Payload(func() {
			Member("alias", String, "Alias of the filter", func() {
				Example("foo")
			})
			Member("name", String, "Name of the filter", func() {
				Example("fetch")
			})
			Member("props", HashOf(String, Any), "Filter properties", NoExample)
			Member("condition", String, "Conditional expression of the output", func() {
				Example("\"foo\" in Tags")
			})
			Required("alias", "name", "condition")
		})
		Response(Created, FilterResponse)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:id"),
		)
		Description("Update a filter")
		Params(func() {
			Param("id", Integer, "Filter ID")
		})
		Payload(func() {
			Member("alias", String, "Alias of the filter", func() {
				Example("foo")
			})
			Member("props", HashOf(String, Any), "Filter properties", NoExample)
			Member("condition", String, "Conditional expression of the output", func() {
				Example("\"foo\" in Tags")
			})
			Member("enabled", Boolean, "Filter status", NoExample)
		})
		Response(OK, FilterResponse)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:id"),
		)
		Description("Delete a filter")
		Params(func() {
			Param("id", Integer, "Filter ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

// FilterResponse is the filter resource media type.
var FilterResponse = MediaType("application/vnd.feedpushr.filter.v1+json", func() {
	Description("A filter")
	TypeName("FilterResponse")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Integer, "ID of the filter", func() {
			Example(1)
		})
		Attribute("alias", String, "Alias of the filter", func() {
			Example("foo")
		})
		Attribute("name", String, "Name of the filter", func() {
			Example("fetch")
		})
		Attribute("desc", String, "Description of the filter", func() {
			Example("This filter will...")
		})
		Attribute("props", HashOf(String, Any), "Filter properties", NoExample)
		Attribute("condition", String, "Conditional expression of the filter", func() {
			Example("\"foo\" in Tags")
		})
		Attribute("enabled", Boolean, "Status", func() {
			Default(false)
		})
		Required("id", "alias", "name", "desc", "condition")
	})

	View("default", func() {
		Attribute("id")
		Attribute("alias")
		Attribute("name")
		Attribute("desc")
		Attribute("props")
		Attribute("condition")
		Attribute("enabled")
	})
})

// FilterSpecResponse is the filter specification media type.
var FilterSpecResponse = MediaType("application/vnd.feedpushr.filter-spec.v1+json", func() {
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
		Attribute("props", CollectionOf("application/vnd.feedpushr.prop-spec.v1+json"))
		Required("name", "desc", "props")
	})

	View("default", func() {
		Attribute("name")
		Attribute("desc")
		Attribute("props")
	})
})
