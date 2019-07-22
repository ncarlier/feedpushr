package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("feed", func() {

	DefaultMedia(Feed)
	BasePath("/feeds")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all feeds")
		Params(func() {
			Param("page", Integer, "Page to fetch", func() {
				Default(1)
				Minimum(1)
				Example(5)
			})
			Param("limit", Integer, "Fetch limit", func() {
				Default(10)
				Minimum(1)
				Example(10)
			})
		})
		Response(OK, func() {
			Media(CollectionOf(Feed, func() {
				View("default")
				View("tiny")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("get", func() {
		Routing(
			GET("/:id"),
		)
		Description("Retrieve feed with given ID")
		Params(func() {
			Param("id", String, "Feed ID")
		})
		Response(OK, func() {
			Media(Feed)
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create a new feed")
		Params(func() {
			Param("url", String, "Feed URL", func() {
				Example("http://www.hashicorp.com/feed.xml")
				Format("uri")
			})
			Param("title", String, "Feed title (will overide official feed title)", func() {
				Example("A cool website")
			})
			Param("tags", String, "Comma separated list of tags", func() {
				Example("foo,bar")
			})
			Required("url")
		})
		Response(Created, Feed)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:id"),
		)
		Description("Update a feed")
		Params(func() {
			Param("id", String, "Feed ID")
			Param("title", String, "Feed title", func() {
				Example("A cool website")
			})
			Param("tags", String, "Comma separated list of tags", func() {
				Example("foo,bar")
			})
		})
		Response(OK, Feed)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:id"),
		)
		Description("Delete a feed")
		Params(func() {
			Param("id", String, "Feed ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("start", func() {
		Routing(
			POST("/:id/start"),
		)
		Description("Start feed aggregation")
		Response(Accepted)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("stop", func() {
		Routing(
			POST("/:id/stop"),
		)
		Description("Stop feed aggregation")
		Response(Accepted)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("filter", func() {
	DefaultMedia(Filter)
	BasePath("/filters")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all filters definitions")
		Response(OK, func() {
			Media(CollectionOf(Filter, func() {
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
			Media(CollectionOf(FilterSpec, func() {
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
		Response(OK, Filter)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create a new filter")
		Payload(func() {
			Member("name", String, "Name of the filter", func() {
				Example("fetch")
			})
			Member("props", HashOf(String, Any), "Filter properties", NoExample)
			Member("tags", String, "Comma separated list of tags", func() {
				Example("foo,bar")
			})
			Required("name")
		})
		Response(Created, Filter)
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
			Member("props", HashOf(String, Any), "Filter properties", NoExample)
			Member("tags", String, "Comma separated list of tags", func() {
				Example("foo,bar")
			})
			Member("enabled", Boolean, "Filter status", NoExample)
		})
		Response(OK, Filter)
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

var _ = Resource("output", func() {
	DefaultMedia(Output)
	BasePath("/outputs")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all outputs definitions")
		Response(OK, func() {
			Media(CollectionOf(Output, func() {
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
			Media(CollectionOf(OutputSpec, func() {
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
		Response(OK, Output)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create a new output")
		Payload(func() {
			Member("name", String, "Name of the output", func() {
				Example("http")
			})
			Member("props", HashOf(String, Any), "Output properties", NoExample)
			Member("tags", String, "Comma separated list of tags", func() {
				Example("foo,bar")
			})
			Required("name")
		})
		Response(Created, Output)
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
			Member("props", HashOf(String, Any), "Output properties", NoExample)
			Member("tags", String, "Comma separated list of tags", func() {
				Example("foo,bar")
			})
			Member("enabled", Boolean, "Output status", NoExample)
		})
		Response(OK, Output)
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

var _ = Resource("pshb", func() {
	BasePath("/pshb")

	Action("sub", func() {
		Routing(
			GET(""),
		)
		Description("Callback to validate the (un)subscription to the topic of a Hub")
		Params(func() {
			Param("hub.topic", String, "The topic URL given in the corresponding subscription request", func() {
				Format("uri")
			})
			Param("hub.mode", String, "The literal string \"subscribe\" or \"unsubscribe\"", func() {
				Enum("subscribe", "unsubscribe")
			})
			Param("hub.challenge", String, "A hub-generated random string")
			Param("hub.lease_seconds", Integer, "The hub-determined number of seconds that the subscription will stay active before expiring")
			Required("hub.topic", "hub.mode", "hub.challenge")
		})
		Response(OK)
		Response(BadRequest, ErrorMedia)
	})

	Action("pub", func() {
		Routing(
			POST(""),
		)
		Description("Publication endpoint for PSHB hubs")
		Response(OK)
		Response(BadRequest, ErrorMedia)
	})
})
