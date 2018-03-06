package controller

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/goadesign/goa"
	"github.com/mmcdole/gofeed"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/aggregator"
	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/output"
	"github.com/ncarlier/feedpushr/pkg/store"
)

// PshbController implements the pshb resource.
type PshbController struct {
	*goa.Controller
	db         store.DB
	parser     *gofeed.Parser
	output     *output.Manager
	aggregator *aggregator.Manager
}

// NewPshbController creates a pshb controller.
func NewPshbController(service *goa.Service, db store.DB, am *aggregator.Manager, om *output.Manager) *PshbController {
	return &PshbController{
		Controller: service.NewController("PshbController"),
		db:         db,
		output:     om,
		aggregator: am,
		parser:     gofeed.NewParser(),
	}
}

// Pub is the Hub callback to send topic updates.
func (c *PshbController) Pub(ctx *app.PubPshbContext) error {
	body, err := common.GetNormalizedBody(ctx.Response)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	feed, err := c.parser.Parse(body)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	c.output.Send(builder.NewArticles(feed.Items))

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
		c.aggregator.RestartFeedAggregator(id, delay)
	} else if ctx.HubMode == "unsubscribe" {
		// Notify the aggregator to resume
		c.aggregator.RestartFeedAggregator(id, 0)
	}

	return ctx.OK([]byte(ctx.HubChallenge))
}
