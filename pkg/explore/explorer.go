package explore

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

// SearchResult is the result of a feed search
type SearchResult struct {
	Title   string `json:"title,omitempty"`
	Desc    string `json:"desc,omitempty"`
	XMLURL  string `json:"xmlUrl,omitempty"`
	HTMLURL string `json:"htmlUrl,omitempty"`
}

// SearchResults is a list of search result
type SearchResults []SearchResult

// Explorer is the interface of a feed explorer
type Explorer interface {
	Search(q string) (*SearchResults, error)
}

// NewExplorer create new explorer for a given provider
func NewExplorer(provider string) (Explorer, error) {
	var explorer Explorer

	switch provider {
	case "rsssearchhub", "default":
		explorer = NewRSSSearchHubExplorerProvider()
		log.Info().Str("component", "explorer").Str("provider", provider).Msg("feed explorer configured")
	default:
		return nil, fmt.Errorf("unsupported explorer provider: %s", provider)
	}
	return explorer, nil
}

func searchByURL(u url.URL) (*SearchResults, error) {
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

	results := SearchResults{}
	for _, u := range urls {
		results = append(results, SearchResult{
			XMLURL: u,
		})
	}

	return &results, nil
}
