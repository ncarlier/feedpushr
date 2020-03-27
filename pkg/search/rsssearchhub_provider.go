package search

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// RSSSearchHubSearchProvider use RSS Search HUB as search engine provider
type RSSSearchHubSearchProvider struct {
}

// NewRSSSearchHubSearchProvider create new search engine using RSS Search HUB
func NewRSSSearchHubSearchProvider() Engine {
	return &RSSSearchHubSearchProvider{}
}

// Search search feeds from a query
func (p *RSSSearchHubSearchProvider) Search(q string) (*Results, error) {
	if u, err := url.ParseRequestURI(q); err == nil {
		if u.Scheme == "" {
			u.Scheme = "http"
		}
		return searchByURL(*u)
	}
	return p.searchByQuery(q)
}

func (p *RSSSearchHubSearchProvider) searchByQuery(q string) (*Results, error) {
	doc, err := goquery.NewDocument("https://www.rsssearchhub.com/feeds?q=" + url.QueryEscape(q))
	if err != nil {
		return nil, errors.Wrapf(err, "querying RSS search HUB with %s", q)
	}

	results := Results{}
	doc.Find(".___rsssearchhubfeed .result").Each(func(i int, s *goquery.Selection) {
		result := Result{}
		result.Title = s.Find("h2 > a").Text()
		result.Text = s.Find("p.description").Text()
		result.HTMLURL = s.Find("p.url > span").Text()
		result.XMLURL = s.Find("p.url > a").Text()
		results = append(results, result)
	})
	return &results, nil
}
