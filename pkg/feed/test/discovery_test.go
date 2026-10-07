package builder_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ncarlier/feedpushr/v3/pkg/feed"
	"github.com/stretchr/testify/require"
)

func TestExtractFeedLinksReferences(t *testing.T) {
	tests := []struct {
		name string
		href string
		want string
	}{
		{"file", "feed.xml", "https://example.com/blog/feed.xml"},
		{"parent", "../feed.xml", "https://example.com/feed.xml"},
		{"dot", "./feed.xml", "https://example.com/blog/feed.xml"},
		{"query", "?feed=rss", "https://example.com/blog/index.html?feed=rss"},
		{"root", "/feed.xml", "https://example.com/feed.xml"},
		{"scheme relative", "//feeds.example.com/feed.xml", "https://feeds.example.com/feed.xml"},
		{"absolute", "https://feeds.example.com/feed.xml", "https://feeds.example.com/feed.xml"},
		{"absolute unchanged", "https://feeds.example.com/a/../feed.xml", "https://feeds.example.com/a/../feed.xml"},
		{"escaped path", "feed%20name.xml", "https://example.com/blog/feed%20name.xml"},
		{"fragment", "feed.xml#items", "https://example.com/blog/feed.xml#items"},
		{"invalid escape", "feed%zz.xml", ""},
		{"empty", "", ""},
		{"document fragment", "#items", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			page := fmt.Sprintf(`<html><head><link rel="alternate" type="application/rss+xml" href="%s"></head><body></body></html>`, tt.href)
			links, err := feed.ExtractFeedLinks(strings.NewReader(page), "https://example.com/blog/index.html")
			require.NoError(t, err)
			if tt.want == "" {
				require.Empty(t, links)
				return
			}
			require.Equal(t, []string{tt.want}, links)
		})
	}
}

func TestNewFeedRelativeDiscovery(t *testing.T) {
	tests := []struct {
		name     string
		href     string
		feedPath string
	}{
		{"file", "feed.xml", "/blog/feed.xml"},
		{"parent", "../feed.xml", "/feed.xml"},
		{"query", "?feed=rss", "/blog/index.html?feed=rss"},
		{"root", "/feed.xml", "/feed.xml"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.RequestURI() {
				case tt.feedPath:
					w.Header().Set("Content-Type", "application/rss+xml")
					fmt.Fprint(w, `<rss version="2.0"><channel><title>Local feed</title><link>https://example.com/</link><description>Test feed</description></channel></rss>`)
				case "/blog/index.html":
					w.Header().Set("Content-Type", "text/html")
					fmt.Fprintf(w, `<html><head><link type="application/rss+xml" href="%s"></head><body></body></html>`, tt.href)
				default:
					http.NotFound(w, r)
				}
			}))
			defer server.Close()

			result, err := feed.NewFeed(server.URL+"/blog/index.html", nil)
			require.NoError(t, err)
			require.Equal(t, server.URL+tt.feedPath, result.XMLURL)
			require.Equal(t, feed.GetFeedID(server.URL+tt.feedPath), result.ID)
			require.Equal(t, "Local feed", result.Title)
		})
	}
}

func TestNewFeedDiscoveryAfterRedirect(t *testing.T) {
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/blog/":
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, `<html><head><link type="application/rss+xml" href="feed.xml"></head><body></body></html>`)
		case "/blog/feed.xml":
			w.Header().Set("Content-Type", "application/rss+xml")
			fmt.Fprint(w, `<rss version="2.0"><channel><title>Redirected feed</title><link>https://example.com/</link><description>Test feed</description></channel></rss>`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer target.Close()
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL+"/blog/", http.StatusFound)
	}))
	defer origin.Close()

	result, err := feed.NewFeed(origin.URL+"/old", nil)
	require.NoError(t, err)
	require.Equal(t, target.URL+"/blog/feed.xml", result.XMLURL)
	require.Equal(t, "Redirected feed", result.Title)
}
