package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

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
