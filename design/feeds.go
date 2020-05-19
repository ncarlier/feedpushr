package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("feed", func() {

	DefaultMedia(FeedResponse)
	BasePath("/feeds")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all feeds")
		Params(func() {
			Param("q", String, "Search query")
			Param("page", Integer, "Page to fetch", func() {
				Default(1)
				Minimum(1)
				Example(5)
			})
			Param("size", Integer, "Page size", func() {
				Default(10)
				Minimum(1)
				Example(10)
			})
		})
		Response(OK, func() {
			Media(FeedsPageResponse)
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
			Media(FeedResponse)
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
			Param("enable", Boolean, "Feed activation status", func() {
				Example(true)
			})
			Required("url")
		})
		Response(Created, FeedResponse)
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
		Response(OK, FeedResponse)
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

// FeedResponse is the feed resource media type.
var FeedResponse = MediaType("application/vnd.feedpushr.feed.v2+json", func() {
	Description("A RSS feed")
	TypeName("FeedResponse")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", String, "ID of feed (MD5 of the xmlUrl)", func() {
			Example("5bfb841c028281c0051828c115fd1f50")
		})
		Attribute("xmlUrl", String, "URL of the XML feed", func() {
			Example("http://www.hashicorp.com/feed.xml")
		})
		Attribute("htmlUrl", String, "URL of the feed website", func() {
			Example("http://www.hashicorp.com/blog")
		})
		Attribute("hubUrl", String, "URL of the PubSubHubbud hub", func() {
			Example("http://pubsubhubbub.appspot.com")
		})
		Attribute("title", String, "Title of the Feed", func() {
			Example("Hashicorp Blog")
		})
		Attribute("tags", ArrayOf(String), "List of tags", func() {
			Example([]string{"foo", "bar"})
		})
		Attribute("status", String, "Aggregation status", func() {
			Enum("running", "stopped")
		})
		Attribute("lastCheck", DateTime, "Last aggregation pass")
		Attribute("nextCheck", DateTime, "Next aggregation pass")
		Attribute("errorMsg", String, "Last aggregation error")
		Attribute("errorCount", Integer, "Number of consecutive aggregation errors")
		Attribute("nbProcessedItems", Integer, "Total number of processed items")
		Attribute("cdate", DateTime, "Date of creation")
		Attribute("mdate", DateTime, "Date of modification")

		Required("id", "xmlUrl", "title", "cdate", "mdate")
	})

	View("default", func() {
		Attribute("id")
		Attribute("xmlUrl")
		Attribute("htmlUrl")
		Attribute("hubUrl")
		Attribute("title")
		Attribute("tags")
		Attribute("status")
		Attribute("lastCheck")
		Attribute("nextCheck")
		Attribute("errorMsg")
		Attribute("errorCount")
		Attribute("nbProcessedItems")
		Attribute("cdate")
		Attribute("mdate")
	})

	View("tiny", func() {
		Description("tiny is the view used to list feeds")
		Attribute("id")
		Attribute("xmlUrl")
		Attribute("title")
		Attribute("tags")
		Attribute("cdate")
	})

	View("link", func() {
		Attribute("id")
		Attribute("xmlUrl")
	})
})

// FeedsPageResponse is the feeds page resource media type.
var FeedsPageResponse = MediaType("application/vnd.feedpushr.feeds-page.v2+json", func() {
	Description("A pagignated list of feeds")
	TypeName("FeedsPageResponse")
	ContentType("application/json")
	Attributes(func() {
		Attribute("total", Integer, "Total number of feeds", func() {
			Example(99)
		})
		Attribute("current", Integer, "Current page number", func() {
			Example(1)
		})
		Attribute("size", Integer, "Max number of feeds by page", func() {
			Example(100)
		})
		Attribute("data", CollectionOf(FeedResponse), "List of feeds")
		Required("total", "current", "size", "data")
	})

	View("default", func() {
		Attribute("total")
		Attribute("current")
		Attribute("size")
		Attribute("data")
	})
})
