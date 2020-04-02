package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("explore", func() {
	DefaultMedia(FilterResponse)
	BasePath("/explore")

	Action("get", func() {
		Routing(
			GET(""),
		)
		Description("Search RSS feed using external search engine")
		Params(func() {
			Param("q", String, "Search query", func() {
				Example("tech blog")
			})
		})
		Response(OK, func() {
			Media(CollectionOf(ExploreResponse, func() {
				View("default")
			}))
		})
		Response(BadRequest, ErrorMedia)
	})
})

// ExploreResponse is the explore result specification media type.
var ExploreResponse = MediaType("application/vnd.feedpushr.explore.v2+json", func() {
	Description("The search result")
	TypeName("ExploreResponse")
	ContentType("application/json")
	Attributes(func() {
		Attribute("title", String, "Feed title", func() {
			Example("Blog news...")
		})
		Attribute("desc", String, "Feed description", func() {
			Example("A short description...")
		})
		Attribute("xmlUrl", String, "URL of the XML feed", func() {
			Example("http://www.hashicorp.com/feed.xml")
		})
		Attribute("htmlUrl", String, "URL of the feed website", func() {
			Example("http://www.hashicorp.com/blog")
		})
		Required("title", "desc", "xmlUrl", "htmlUrl")
	})

	View("default", func() {
		Attribute("title")
		Attribute("desc")
		Attribute("xmlUrl")
		Attribute("htmlUrl")
	})
})
