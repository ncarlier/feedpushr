package controller

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/goadesign/goa"
	"github.com/mmcdole/gofeed"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
	"github.com/ncarlier/feedpushr/v3/pkg/aggregator"
	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/feed"
	"github.com/ncarlier/feedpushr/v3/pkg/helper"
	"github.com/ncarlier/feedpushr/v3/pkg/output"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var maxSubscriptionTTL = time.Duration(72) * time.Hour

// PshbController implements the pshb resource.
type PshbController struct {
	*goa.Controller
	db         store.DB
	parser     *gofeed.Parser
	outputs    *output.Manager
	aggregator *aggregator.Manager
	log        zerolog.Logger
}

// NewPshbController creates a pshb controller.
func NewPshbController(service *goa.Service, db store.DB, am *aggregator.Manager, outputs *output.Manager) *PshbController {
	parser := gofeed.NewParser()
	parser.AtomTranslator = feed.NewCustomAtomTranslator()
	parser.RSSTranslator = feed.NewCustomRSSTranslator()
	return &PshbController{
		Controller: service.NewController("PshbController"),
		db:         db,
		outputs:    outputs,
		aggregator: am,
		parser:     parser,
		log:        log.With().Str("component", "pshb-ctrl").Logger(),
	}
}

// Pub is the Hub callback to send topic updates.
func (c *PshbController) Pub(ctx *app.PubPshbContext) error {
	body, err := helper.GetNormalizedBodyFromRequest(ctx.Request)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	parsedFeed, err := c.parser.Parse(body)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	link := parsedFeed.FeedLink
	if self, ok := parsedFeed.Custom["self"]; ok {
		link = self
	}

	id := feed.GetFeedID(link)
	_feed, err := c.db.GetFeed(id)
	if err != nil {
		c.log.Warn().Str("id", id).Str("link", link).Msg("PSHB callback received an unknown feed link")
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	c.outputs.Push(feed.NewArticles(_feed, parsedFeed.Items))

	return ctx.OK([]byte("ok"))
}

// Sub is the Hub callback to validate the (un)subscription.
func (c *PshbController) Sub(ctx *app.SubPshbContext) error {
	// Compute feed ID
	hasher := md5.New()
	hasher.Write([]byte(ctx.HubTopic))
	id := hex.EncodeToString(hasher.Sum(nil))

	// Get feed from DB
	_, err := c.db.GetFeed(id)
	if err != nil {
		// Acknowledge the unsubscribtion
		if err == common.ErrFeedNotFound && ctx.HubMode == "unsubscribe" {
			return ctx.OK([]byte(ctx.HubChallenge))
		}
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	if ctx.HubMode == "subscribe" && ctx.HubLeaseSeconds != nil {
		// Notify the aggregator to wait until the lease is over
		delay := time.Duration(*ctx.HubLeaseSeconds) * time.Second
		if delay > maxSubscriptionTTL {
			delay = maxSubscriptionTTL
		}
		c.aggregator.RestartFeedAggregator(id, delay)
		c.log.Info().Str("id", id).Msg("PSHB subscription activated")
	} else if ctx.HubMode == "unsubscribe" {
		// Notify the aggregator to resume
		c.aggregator.RestartFeedAggregator(id, 0)
		c.log.Info().Str("id", id).Msg("PSHB subscription deactivated")
	}

	return ctx.OK([]byte(ctx.HubChallenge))
}
