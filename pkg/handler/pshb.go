package handler

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/ncarlier/feedpushr/v3/pkg/api"
	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/feed"
	"github.com/ncarlier/feedpushr/v3/pkg/helper"
)

var maxSubscriptionTTL = time.Duration(72) * time.Hour

// PshbPublish handles POST /v2/pshb (Hub callback to send topic updates)
func (h *Handlers) PshbPublish(w http.ResponseWriter, r *http.Request) {
	res := api.NewResponse(w)

	body, err := helper.GetNormalizedBodyFromRequest(r)
	if err != nil {
		res.BadRequest(err)
		return
	}

	parser := gofeed.NewParser()
	parser.AtomTranslator = feed.NewCustomAtomTranslator()
	parser.RSSTranslator = feed.NewCustomRSSTranslator()

	parsedFeed, err := parser.Parse(body)
	if err != nil {
		res.BadRequest(err)
		return
	}

	link := parsedFeed.FeedLink
	if self, ok := parsedFeed.Custom["self"]; ok {
		link = self
	}

	id := feed.GetFeedID(link)
	_feed, err := h.db.GetFeed(r.Context(), id)
	if err != nil {
		h.log.Warn().Str("id", id).Str("link", link).Msg("PSHB callback received an unknown feed link")
		res.BadRequest(err)
		return
	}

	h.outputs.Push(feed.NewArticles(_feed, parsedFeed.Items))
	res.Raw("text/plain", []byte("ok"))
}

// PshbSubscribe handles GET /v2/pshb (Hub callback to validate the subscription)
func (h *Handlers) PshbSubscribe(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	hubMode := req.QueryParam("hub.mode")
	hubTopic := req.QueryParam("hub.topic")
	hubChallenge := req.QueryParam("hub.challenge")
	hubLeaseSeconds := req.QueryParamInt("hub.lease_seconds", 0)

	// Compute feed ID
	hasher := md5.New()
	hasher.Write([]byte(hubTopic))
	id := hex.EncodeToString(hasher.Sum(nil))

	// Get feed from DB
	_, err := h.db.GetFeed(r.Context(), id)
	if err != nil {
		// Acknowledge the unsubscription
		if err == common.ErrFeedNotFound && hubMode == "unsubscribe" {
			res.Raw("text/plain", []byte(hubChallenge))
			return
		}
		res.BadRequest(err)
		return
	}

	if hubMode == "subscribe" && hubLeaseSeconds > 0 {
		// Notify the aggregator to wait until the lease is over
		delay := time.Duration(hubLeaseSeconds) * time.Second
		if delay > maxSubscriptionTTL {
			delay = maxSubscriptionTTL
		}
		h.aggregator.RestartFeedAggregator(id, delay)
		h.log.Info().Str("id", id).Msg("PSHB subscription activated")
	} else if hubMode == "unsubscribe" {
		// Notify the aggregator to resume
		h.aggregator.RestartFeedAggregator(id, 0)
		h.log.Info().Str("id", id).Msg("PSHB subscription deactivated")
	}

	res.Raw("text/plain", []byte(hubChallenge))
}
