package pshb_test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/pshb"
)

func TestGetSubscriptionDetailsURL(t *testing.T) {
	hub := "https://pubsubhubbub.appspot.com"
	callback := "https://reader.nunux.org/pubsubhubbud/callback"
	topic := "http://feeds.feedburner.com/test"
	u := pshb.GetSubscriptionDetailsURL(hub, topic, callback)
	expected := "https://pubsubhubbub.appspot.com/subscription-details?hub.callback=https%3A%2F%2Freader.nunux.org%2Fpubsubhubbud%2Fcallback&hub.topic=http%3A%2F%2Ffeeds.feedburner.com%2Ftest"
	assert.Equal(t, expected, *u, "Bad PSHB subscription URL")
}
