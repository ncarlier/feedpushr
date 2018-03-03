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
				Maximum(100)
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
		Description("Retrieve feed with given id")
		Params(func() {
			Param("id", String, "Feed ID")
		})
		Response(OK)
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
			Required("url")
		})
		Response(Created, "/feeds/[0-9a-f]+")
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

var _ = Resource("opml", func() {

	BasePath("/opml")

	Action("get", func() {
		Routing(
			GET(""),
		)
		Description("Get all feeds as an OMPL format")
		Response(OK)
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
	Files("/swagger.json", "var/public/swagger.json")
})

var _ = Resource("vars", func() {
	BasePath("/vars")

	Action("get", func() {
		Routing(
			GET(""),
		)
		Description("Get all internals exp vars")
		Response(OK)
		Response(BadRequest, ErrorMedia)
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
