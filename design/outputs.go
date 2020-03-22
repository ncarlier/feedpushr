package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("output", func() {
	DefaultMedia(OutputResponse)
	BasePath("/outputs")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all outputs definitions")
		Response(OK, func() {
			Media(CollectionOf(OutputResponse, func() {
				View("default")
			}))
		})
	})

	Action("specs", func() {
		Routing(
			GET("/_specs"),
		)
		Description("Retrieve all output types available")
		Response(OK, func() {
			Media(CollectionOf(OutputSpecResponse, func() {
				View("default")
			}))
		})
	})

	Action("get", func() {
		Routing(
			GET("/:id"),
		)
		Description("Retrieve output with given ID")
		Params(func() {
			Param("id", Integer, "Output ID")
		})
		Response(OK, OutputResponse)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create a new output")
		Payload(func() {
			Member("alias", String, "Alias of the output", func() {
				Example("foo")
			})
			Member("name", String, "Name of the output", func() {
				Example("http")
			})
			Member("props", HashOf(String, Any), "Output properties", NoExample)
			Member("condition", String, "Conditional expression of the output", func() {
				Example("\"foo\" in Tags")
			})
			Required("alias", "name", "condition")
		})
		Response(Created, OutputResponse)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:id"),
		)
		Description("Update an output")
		Params(func() {
			Param("id", Integer, "Output ID")
		})
		Payload(func() {
			Member("alias", String, "Alias of the output", func() {
				Example("foo")
			})
			Member("props", HashOf(String, Any), "Output properties", NoExample)
			Member("condition", String, "Conditional expression of the output", func() {
				Example("\"foo\" in Tags")
			})
			Member("enabled", Boolean, "Output status", NoExample)
		})
		Response(OK, OutputResponse)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:id"),
		)
		Description("Delete an output")
		Params(func() {
			Param("id", Integer, "Output ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

// OutputResponse is the output resource media type.
var OutputResponse = MediaType("application/vnd.feedpushr.output.v1+json", func() {
	Description("The output channel")
	TypeName("OutputResponse")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Integer, "ID of the output", func() {
			Example(1)
		})
		Attribute("alias", String, "Alias of the output channel", func() {
			Example("foo")
		})
		Attribute("name", String, "Name of the output channel", func() {
			Example("fetch")
		})
		Attribute("desc", String, "Description of the output channel", func() {
			Example("New articles are sent as JSON document to...")
		})
		Attribute("props", HashOf(String, Any), "Output channel properties", NoExample)
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
		Attribute("name")
		Attribute("alias")
		Attribute("desc")
		Attribute("props")
		Attribute("condition")
		Attribute("enabled")
	})
})

// OutputSpecResponse is the output specification media type.
var OutputSpecResponse = MediaType("application/vnd.feedpushr.output-spec.v1+json", func() {
	Description("The output channel specification")
	TypeName("OutputSpecResponse")
	ContentType("application/json")
	Attributes(func() {
		Attribute("name", String, "Name of the output channel", func() {
			Example("fetch")
		})
		Attribute("desc", String, "Description of the output channel", func() {
			Example("New articles are sent as JSON document to...")
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

// PropSpecResponse is the property specification media type.
var PropSpecResponse = MediaType("application/vnd.feedpushr.prop-spec.v1+json", func() {
	Description("The specification of a property")
	TypeName("PropSpec")
	ContentType("application/json")
	Attributes(func() {
		Attribute("name", String, "Name of the property", func() {
			Example("url")
		})
		Attribute("desc", String, "Description of the output channel", func() {
			Example("New articles are sent as JSON document to...")
		})
		Attribute("type", String, "Property type ('text', 'url', ...)", func() {
			Example("text")
		})
		Attribute("options", HashOf(String, String), "Property options")
		Required("name", "desc", "type")
	})

	View("default", func() {
		Attribute("name")
		Attribute("desc")
		Attribute("type")
		Attribute("options")
	})
})
