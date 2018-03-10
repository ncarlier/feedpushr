package builder

import (
	"fmt"

	"github.com/mmcdole/gofeed"
	"github.com/mmcdole/gofeed/atom"
	ext "github.com/mmcdole/gofeed/extensions"
	"github.com/mmcdole/gofeed/rss"
)

// CustomAtomTranslator is a custom GoFeed Atom translator created to extract Hub link.
type CustomAtomTranslator struct {
	defaultTranslator *gofeed.DefaultAtomTranslator
}

// NewCustomAtomTranslator creates nes custom GoFeed Atom translator.
func NewCustomAtomTranslator() *CustomAtomTranslator {
	t := &CustomAtomTranslator{}

	t.defaultTranslator = &gofeed.DefaultAtomTranslator{}
	return t
}

// Translate Atom feed into generic feed and extract Hub link if present.
func (ct *CustomAtomTranslator) Translate(feed interface{}) (*gofeed.Feed, error) {
	rss, found := feed.(*atom.Feed)
	if !found {
		return nil, fmt.Errorf("Feed did not match expected type of *atom.Feed")
	}

	f, err := ct.defaultTranslator.Translate(rss)
	if err != nil {
		return nil, err
	}

	hub := firstLinkWithType("hub", rss.Links)
	if hub != nil {
		if f.Custom == nil {
			f.Custom = make(map[string]string)
		}
		f.Custom["hub"] = hub.Href
	}

	return f, nil
}

// CustomRSSTranslator is a custom GoFeed RSS translator created to extract Hub link.
type CustomRSSTranslator struct {
	defaultTranslator *gofeed.DefaultRSSTranslator
}

// NewCustomRSSTranslator creates nes custom GoFeed RSS translator.
func NewCustomRSSTranslator() *CustomRSSTranslator {
	t := &CustomRSSTranslator{}

	t.defaultTranslator = &gofeed.DefaultRSSTranslator{}
	return t
}

// Translate RSS feed into generic feed and extract Hub link if present.
func (ct *CustomRSSTranslator) Translate(feed interface{}) (*gofeed.Feed, error) {
	rss, found := feed.(*rss.Feed)
	if !found {
		return nil, fmt.Errorf("Feed did not match expected type of *rss.Feed")
	}

	f, err := ct.defaultTranslator.Translate(rss)
	if err != nil {
		return nil, err
	}

	var hub string
	atomExtensions := extensionsForKeys([]string{"atom", "atom10", "atom03"}, rss.Extensions)
	for _, ex := range atomExtensions {
		if links, ok := ex["link"]; ok {
			for _, l := range links {
				if l.Attrs["rel"] == "hub" {
					hub = l.Attrs["href"]
				}
			}
		}
	}

	if hub != "" {
		if f.Custom == nil {
			f.Custom = make(map[string]string)
		}
		f.Custom["hub"] = hub
	}

	return f, nil
}

func firstLinkWithType(linkType string, links []*atom.Link) *atom.Link {
	if links == nil {
		return nil
	}

	for _, link := range links {
		if link.Rel == linkType {
			return link
		}
	}
	return nil
}

func extensionsForKeys(keys []string, extensions ext.Extensions) (matches []map[string][]ext.Extension) {
	matches = []map[string][]ext.Extension{}

	if extensions == nil {
		return
	}

	for _, key := range keys {
		if match, ok := extensions[key]; ok {
			matches = append(matches, match)
		}
	}
	return
}
