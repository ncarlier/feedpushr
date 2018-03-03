package controller

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/aggregator"
	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/store"
)

// FeedController implements the feed resource.
type FeedController struct {
	*goa.Controller
	db         store.DB
	aggregator *aggregator.Manager
}

// NewFeedController creates a feed controller.
func NewFeedController(service *goa.Service, db store.DB, am *aggregator.Manager) *FeedController {
	return &FeedController{
		Controller: service.NewController("FeedController"),
		db:         db,
		aggregator: am,
	}
}

// Create creates a new feed
func (c *FeedController) Create(ctx *app.CreateFeedContext) error {
	feed, err := builder.NewFeed(ctx.URL)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	err = c.db.SaveFeed(feed)
	if err != nil {
		return goa.ErrInternal(err)
	}
	fa := c.aggregator.RegisterFeedAggregator(feed)
	fa.Start()
	return ctx.Created()
}

// Delete removes a feed
func (c *FeedController) Delete(ctx *app.DeleteFeedContext) error {
	c.aggregator.UnRegisterFeedAggregator(ctx.ID)
	_, err := c.db.DeleteFeed(ctx.ID)
	if err != nil {
		if err == common.ErrFeedNotFound {
			return ctx.NotFound()
		}
		return goa.ErrInternal(err)
	}
	return ctx.NoContent()
}

// List shows all feeds
func (c *FeedController) List(ctx *app.ListFeedContext) error {
	feeds, err := c.db.ListFeeds(ctx.Page, ctx.Limit)
	if err != nil {
		return goa.ErrInternal(err)
	}

	for i := 0; i < len(*feeds); i++ {
		feed := (*feeds)[i]
		// Get feed details with aggregation status
		fa := c.aggregator.GetFeedAggregator(feed.ID)
		if fa != nil {
			feed = fa.GetFeedWithAggregationStatus()
		}
		(*feeds)[i] = feed
	}
	return ctx.OK(*feeds)
}

// Get shows a feed
func (c *FeedController) Get(ctx *app.GetFeedContext) error {
	// Get feed from the database
	feed, err := c.db.GetFeed(ctx.ID)
	if err != nil {
		if err == common.ErrFeedNotFound {
			return ctx.NotFound()
		}
		return goa.ErrInternal(err)
	}
	// Get feed details with aggregation status
	fa := c.aggregator.GetFeedAggregator(feed.ID)
	if fa != nil {
		feed = fa.GetFeedWithAggregationStatus()
	}
	return ctx.OK(feed)
}

// Start starts feed aggregation
func (c *FeedController) Start(ctx *app.StartFeedContext) error {
	feed, err := c.db.GetFeed(ctx.ID)
	if err != nil {
		if err == common.ErrFeedNotFound {
			return ctx.NotFound()
		}
		return goa.ErrInternal(err)
	}
	fa := c.aggregator.GetFeedAggregator(feed.ID)
	if fa == nil {
		return goa.ErrInternal(fmt.Errorf("feed aggregator not registered"))
	}
	fa.Start()
	return ctx.Accepted()
}

// Stop stops feed aggregation
func (c *FeedController) Stop(ctx *app.StopFeedContext) error {
	feed, err := c.db.GetFeed(ctx.ID)
	if err != nil {
		if err == common.ErrFeedNotFound {
			return ctx.NotFound()
		}
		return goa.ErrInternal(err)
	}
	fa := c.aggregator.GetFeedAggregator(feed.ID)
	if fa == nil {
		return goa.ErrInternal(fmt.Errorf("feed aggregator not registered"))
	}
	fa.Stop()
	return ctx.Accepted()
}
