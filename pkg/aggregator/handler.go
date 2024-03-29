package aggregator

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/ncarlier/feedpushr/v3/pkg/feed"
	"github.com/ncarlier/feedpushr/v3/pkg/helper"
	httpc "github.com/ncarlier/feedpushr/v3/pkg/http"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	errCreateRequest = "unable to create request"
	errDoRequest     = "unable to fetch feed (bad request)"
	errTimeout       = "unable to fetch feed (timeout)"
	errEmptyBody     = "unable to fetch feed (empty)"
	errParssingBody  = "unable to fetch feed (bad body)"
	errBadStatus     = "unable to fetch feed (status: %d)"
)

// FeedHandler handles feed refresh
type FeedHandler struct {
	log        zerolog.Logger
	feed       *model.FeedDef
	status     *FeedStatus
	parser     *gofeed.Parser
	httpClient *http.Client
}

// NewFeedHandler create new feed handler
func NewFeedHandler(feed *model.FeedDef, timeout time.Duration) *FeedHandler {
	handler := FeedHandler{
		log:        log.With().Str("handler", feed.ID).Logger(),
		feed:       feed,
		status:     &FeedStatus{},
		parser:     gofeed.NewParser(),
		httpClient: httpc.New(timeout),
	}
	return &handler
}

// Refresh fetch feed items
func (fh *FeedHandler) Refresh() (FeedStatus, []*model.Article) {
	// defer timer.ExecutionTime(fh.log, time.Now(), "refresh")

	var items []*model.Article

	// Update status check date
	fh.status.CheckedAt = time.Now()

	// Create the request
	req, err := http.NewRequest("GET", fh.feed.XMLURL, nil)
	if err != nil {
		fh.log.Error().Err(err).Msg(errCreateRequest)
		fh.status.Err(err)
		return *fh.status, nil
	}

	// Set custom headers
	req.Header.Set("User-Agent", httpc.UserAgent)

	// Add cache headers
	if fh.status.LastModifiedHeader != "" {
		req.Header.Add("If-Modified-Since", fh.status.LastModifiedHeader)
	}
	if fh.status.EtagHeader != "" {
		req.Header.Add("If-None-Match", fh.status.EtagHeader)
	}

	// Do HTTP call
	resp, err := fh.httpClient.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded {
			fh.log.Warn().Err(err).Msg(errDoRequest)
		} else {
			fh.log.Error().Err(err).Msg(errDoRequest)
		}
		fh.status.Err(err)
		return *fh.status, nil
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		// Update cache headers
		fh.status.EtagHeader = resp.Header.Get("Etag")
		fh.status.LastModifiedHeader = resp.Header.Get("Last-Modified")
		fh.status.ExpiresHeader = resp.Header.Get("Expires")
	case 304:
		// Not modified: returns empty result.
		return *fh.status, items
	default:
		fh.log.Info().Err(err).Msgf(errBadStatus, resp.StatusCode)
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
	body, err := helper.GetNormalizedBodyFromResponse(resp)
	if err != nil {
		fh.log.Error().Err(err).Msg(errParssingBody)
		fh.status.Err(err)
		return *fh.status, nil
	}
	f, err := fh.parser.Parse(body)
	if err != nil {
		fh.log.Error().Err(err).Msg(errParssingBody)
		fh.status.Err(err)
		return *fh.status, nil
	}

	// Reset error status
	fh.status.Err(nil)

	// fh.log.Debug().Int("items", len(f.Items)).Msg("f fetched")

	return *fh.status, feed.NewArticles(fh.feed, f.Items)
}

// ResetStatus reset handler status
func (fh *FeedHandler) ResetStatus() {
	fh.status = &FeedStatus{}
}
