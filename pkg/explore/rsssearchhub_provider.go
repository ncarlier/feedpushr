package explore

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// RSSSearchHubExplorerProvider use RSS Search HUB as feed explorer
type RSSSearchHubExplorerProvider struct {
}

// NewRSSSearchHubExplorerProvider create new feed explorer using RSS Search HUB
func NewRSSSearchHubExplorerProvider() Explorer {
	return &RSSSearchHubExplorerProvider{}
}

// Search feeds from a query
func (p *RSSSearchHubExplorerProvider) Search(q string) (*SearchResults, error) {
	if u, err := url.ParseRequestURI(q); err == nil {
		if u.Scheme == "" {
			u.Scheme = "http"
		}
		return searchByURL(*u)
	}
	return p.searchByQuery(q)
}

func (p *RSSSearchHubExplorerProvider) searchByQuery(q string) (*SearchResults, error) {
	doc, err := goquery.NewDocument("https://www.rsssearchhub.com/feeds?q=" + url.QueryEscape(q))
	if err != nil {
		return nil, errors.Wrapf(err, "querying RSS search HUB with %s", q)
	}

	urls := make(map[string]bool)
	results := SearchResults{}
	doc.Find(".___rsssearchhubfeed .result").Each(func(i int, s *goquery.Selection) {
		result := SearchResult{}
		result.Title = s.Find("h2 > a").Text()
		result.Desc = s.Find("p.description").Text()
		// TODO find a better way to prefix scheme
		result.HTMLURL = "http://" + s.Find("p.url > span").Text()
		result.XMLURL = "http://" + s.Find("p.url > a").Text()
		// TODO go deeper for URL that contains `...`
		// Avoid duplicate entries
		if _, exists := urls[result.XMLURL]; !exists {
			results = append(results, result)
			urls[result.XMLURL] = true
		}
	})
	return &results, nil
}
