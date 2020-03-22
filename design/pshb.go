package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

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

// SubscriptionPayload defines the data structure used in the create subscription request body.
var SubscriptionPayload = Type("SubscriptionPayload", func() {
	Attribute("alias", func() {
		MinLength(2)
		Example("Best app ever")
	})
	Attribute("uri", func() {
		MinLength(5)
		Example("https://api:KEY@api.nunux.org/keeper/v2/documents")
	})
})
