package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Feed is the feed resource media type.
var Feed = MediaType("application/vnd.feedpushr.feed.v1+json", func() {
	Description("A RSS feed")
	TypeName("Feed")
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
		Attribute("text", String, "Text attribute of the Feed", func() {
			Example("RSS Feed")
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
		Attribute("text")
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

// Filter is the filter resource media type.
var Filter = MediaType("application/vnd.feedpushr.filter.v1+json", func() {
	Description("A filter")
	TypeName("Filter")
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
		Attribute("tags", ArrayOf(String), "List of tags", func() {
			Example([]string{"foo", "bar"})
		})
		Attribute("enabled", Boolean, "Status", func() {
			Default(false)
		})
		Required("id", "alias", "name", "desc")
	})

	View("default", func() {
		Attribute("id")
		Attribute("alias")
		Attribute("name")
		Attribute("desc")
		Attribute("props")
		Attribute("tags")
		Attribute("enabled")
	})
})

// Output is the output resource media type.
var Output = MediaType("application/vnd.feedpushr.output.v1+json", func() {
	Description("The output channel")
	TypeName("Output")
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
		Attribute("tags", ArrayOf(String), "List of tags", func() {
			Example([]string{"foo", "bar"})
		})
		Attribute("enabled", Boolean, "Status", func() {
			Default(false)
		})
		Required("id", "alias", "name", "desc")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("alias")
		Attribute("desc")
		Attribute("props")
		Attribute("tags")
		Attribute("enabled")
	})
})

// PropSpec is the property specification media type.
var PropSpec = MediaType("application/vnd.feedpushr.prop-spec.v1+json", func() {
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
		Attribute("options", ArrayOf(String), "Property options")
		Required("name", "desc", "type")
	})

	View("default", func() {
		Attribute("name")
		Attribute("desc")
		Attribute("type")
		Attribute("options")
	})
})

// OutputSpec is the output specification media type.
var OutputSpec = MediaType("application/vnd.feedpushr.output-spec.v1+json", func() {
	Description("The output channel specification")
	TypeName("OutputSpec")
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

// FilterSpec is the filter specification media type.
var FilterSpec = MediaType("application/vnd.feedpushr.filter-spec.v1+json", func() {
	Description("The filter specification")
	TypeName("FilterSpec")
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
