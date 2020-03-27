package search

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"net/url"
	"time"

	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/html"
	"github.com/rs/zerolog/log"
)

// Result is the result of a feed search
type Result struct {
	Title   string
	XMLURL  string
	HTMLURL string
	Text    string
}

// Results is a list of search result
type Results []Result

// Engine is the interface of a feed search engine
type Engine interface {
	Search(q string) (*Results, error)
}

// NewSearchEngine the search engine
func NewSearchEngine(provider string) (Engine, error) {
	var engine Engine

	switch provider {
	case "rsssearchhub", "default":
		engine = NewRSSSearchHubSearchProvider()
		log.Info().Str("component", "search").Str("provider", provider).Msg("feed search engine configured")
	default:
		return nil, fmt.Errorf("unsupported search engine provider: %s", provider)
	}
	return engine, nil
}

func searchByURL(u url.URL) (*Results, error) {
	// Set timeout context
	ctx, cancel := context.WithCancel(context.TODO())
	timeout := time.AfterFunc(common.DefaultTimeout, func() {
		cancel()
	})

	// Create the request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", common.UserAgent)

	// Do HTTP call
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	timeout.Stop()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("http error: %s", res.Status)
	}

	// Get content-type
	contentTypeHeader := res.Header.Get("Content-type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		return nil, err
	}

	if contentType != "text/html" {
		return nil, fmt.Errorf("not a valid HTML page: %s", u.String())
	}

	urls, err := html.ExtractFeedLinks(res.Body)
	if err != nil {
		return nil, err
	}
	if len(urls) == 0 {
		return nil, fmt.Errorf("no feed URL found on this page: %s", u.String())
	}

	results := Results{}
	for _, u := range urls {
		results = append(results, Result{
			XMLURL: u,
		})
	}

	return &results, nil
}
