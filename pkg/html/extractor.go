package html

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"golang.org/x/net/html"
)

func extractProperty(token html.Token, prop string) (value string, ok bool) {
	for _, attr := range token.Attr {
		if attr.Key == prop {
			value = attr.Val
			ok = true
		}
	}
	return
}

func absoluteURL(base, path string) (string, error) {
	if strings.HasPrefix(path, "/") {
		baseURL, err := url.Parse(base)
		if err != nil {
			return "", err
		}
		relativeURL, err := url.Parse(path)
		if err != nil {
			return "", err
		}
		u := baseURL.ResolveReference(relativeURL)
		return u.String(), nil
	}
	return path, nil
}

// ExtractFeedLinks extract feed links from HTML content
func ExtractFeedLinks(content io.Reader, base string) (links []string, err error) {
	z := html.NewTokenizer(content)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links, fmt.Errorf("unable to parse HTML content")
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == `body` {
				return links, nil
			}
			if t.Data == "link" {
				linkType, ok := extractProperty(t, "type")
				if ok && common.ValidFeedContentType.MatchString(linkType) {
					if href, ok := extractProperty(t, "href"); ok {
						if link, err := absoluteURL(base, href); err == nil {
							links = append(links, link)
						}
					}
				}
			}
		}
	}
}
