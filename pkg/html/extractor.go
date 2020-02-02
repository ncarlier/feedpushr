package html

import (
	"fmt"
	"io"

	"github.com/ncarlier/feedpushr/v2/pkg/common"
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

// ExtractFeedLinks extract feed links from HTML content
func ExtractFeedLinks(content io.Reader) (links []string, err error) {
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
					href, ok := extractProperty(t, "href")
					if ok {
						links = append(links, href)
					}
				}
			}
		}
	}
}
