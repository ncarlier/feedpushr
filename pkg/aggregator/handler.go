package aggregator

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const defaultTimeout = time.Duration(2) * time.Second

const (
	errCreateRequest = "unable to create request"
	errDoRequest     = "unable to fetch feed (bad request)"
	errEmptyBody     = "unable to fetch feed (empty)"
	errParssingBody  = "ubanle to fetch feed (bad body)"
	errBadStatus     = "unable to fetch feed (status: %d)"
)

// FeedHandler handles feed refresh
type FeedHandler struct {
	log    zerolog.Logger
	feed   *app.Feed
	status *FeedStatus
	parser *gofeed.Parser
}

// NewFeedHandler creats new feed handler
func NewFeedHandler(feed *app.Feed) *FeedHandler {
	handler := FeedHandler{
		log:    log.With().Str("handler", feed.ID).Logger(),
		feed:   feed,
		status: &FeedStatus{},
		parser: gofeed.NewParser(),
	}
	return &handler
}

// Refresh fetch feed items
func (fh *FeedHandler) Refresh() (FeedStatus, []*model.Article) {
	// defer timer.ExecutionTime(fh.log, time.Now(), "refresh")

	var items []*model.Article

	// Set timeout context
	ctx, cancel := context.WithCancel(context.TODO())
	timeout := time.AfterFunc(defaultTimeout, func() {
		cancel()
	})

	// Update status check date
	fh.status.CheckedAt = time.Now()

	// Create the request
	req, err := http.NewRequest("GET", fh.feed.XMLURL, nil)
	if err != nil {
		fh.log.Error().Err(err).Msg(errCreateRequest)
		fh.status.Err(err)
		return *fh.status, nil
	}
	req = req.WithContext(ctx)

	// Add cache headers
	req.Header.Add("If-Modified-Since", fh.status.LastModifiedHeader)
	req.Header.Add("If-None-Match", fh.status.EtagHeader)

	// Do HTTP call
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fh.log.Error().Err(err).Msg(errDoRequest)
		fh.status.Err(err)
		return *fh.status, nil
	}
	defer resp.Body.Close()
	timeout.Stop()

	switch resp.StatusCode {
	case 200, 304:
		// Update cache headers
		fh.status.EtagHeader = resp.Header.Get("Etag")
		fh.status.LastModifiedHeader = resp.Header.Get("Last-Modified")
		fh.status.ExpiresHeader = resp.Header.Get("Expires")
		if resp.StatusCode == 304 {
			// Not modified: returns empty result.
			return *fh.status, items
		}
	default:
		fh.log.Error().Err(err).Msgf(errBadStatus, resp.StatusCode)
		fh.status.Err(fmt.Errorf(errBadStatus, resp.StatusCode))
		return *fh.status, nil
	}

	// Validate that the body is not empty
	if resp.ContentLength == 0 {
		fh.log.Error().Err(err).Msg(errEmptyBody)
		fh.status.Err(fmt.Errorf(errEmptyBody))
		return *fh.status, nil
	}

	// Decode body content
	body, err := common.GetNormalizedBody(resp)
	if err != nil {
		fh.log.Error().Err(err).Msg(errParssingBody)
		fh.status.Err(err)
		return *fh.status, nil
	}
	feed, err := fh.parser.Parse(body)
	if err != nil {
		fh.log.Error().Err(err).Msg(errParssingBody)
		fh.status.Err(err)
		return *fh.status, nil
	}

	// Reset error status
	fh.status.Err(nil)

	// fh.log.Debug().Int("items", len(feed.Items)).Msg("feed fetched")

	return *fh.status, builder.NewArticles(feed.Items)
}
