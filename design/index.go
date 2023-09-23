package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// InfoInfoResponse is the API info media type.
var InfoResponse = MediaType("application/vnd.feedpushr.info.v2+json", func() {
	Description("API info")
	TypeName("Info")
	ContentType("application/json")
	Attributes(func() {
		Attribute("name", String, "Service name", func() {
			Example("feedpushr")
		})
		Attribute("desc", String, "Service description", func() {
			Example("Feed aggregator daemon with sugar on top")
		})
		Attribute("version", String, "Service version", func() {
			Example("v3.0.0")
		})
		Attribute("client_id", String, "Default UI ClientID", func() {
			Example("feedpushr-ui")
		})
		Attribute("_links", HashOf(String, HALLink), "HAL links")
		Required("name", "desc", "version", "client_id", "_links")
	})

	View("default", func() {
		Attribute("name")
		Attribute("desc")
		Attribute("version")
		Attribute("client_id")
		Attribute("_links")
	})
})

// HALLink is the HAL link media type.
var HALLink = MediaType("application/vnd.feedpushr.hal-links.v2+json", func() {
	Description("HAL link")
	TypeName("HALLink")
	ContentType("application/json")
	Attributes(func() {
		Attribute("href", String, "Link's destination", func() {
			Example("url")
		})
		Required("href")
	})

	View("default", func() {
		Attribute("href")
	})
})

var _ = Resource("index", func() {
	Action("get", func() {
		Routing(
			GET("/"),
		)
		Description("Get basic API information.")
		Response(OK, func() {
			Media(InfoResponse)
		})
	})
})
