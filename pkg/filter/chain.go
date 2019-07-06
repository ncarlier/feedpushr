package filter

import (
	"fmt"
	"net/url"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/feedpushr/pkg/plugin"
	"github.com/ncarlier/feedpushr/pkg/store"
)

// Chain contains filter chain
type Chain struct {
	db      store.DB
	filters []model.Filter
}

// LoadChainFilter init chain filter from database
func LoadChainFilter(db store.DB) (*Chain, error) {
	chain := &Chain{
		db: db,
	}
	err := db.ForEachFilter(func(f *app.Filter) error {
		if f == nil {
			return fmt.Errorf("filter is null")
		}
		chain.add(f)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return chain, nil
}

// AddURI add a filter by its URI
func (chain *Chain) AddURI(URI string) error {
	u, err := url.Parse(URI)
	if err != nil {
		return fmt.Errorf("invalid filter URI: %s", URI)
	}
	tags := builder.GetFeedTags(&u.Fragment)
	props := make(map[string]interface{})
	for key, value := range u.Query() {
		props[key] = value[0]
	}
	filter := &app.Filter{
		Name:  u.Scheme,
		Props: props,
		Tags:  tags,
	}
	return chain.Add(filter)
}

// Add a filter into DB then to the current chain filter
func (chain *Chain) Add(filter *app.Filter) error {
	_, err := chain.db.SaveFilter(filter)
	if err != nil {
		return err
	}
	return chain.add(filter)
}

func (chain *Chain) add(filter *app.Filter) error {
	var _filter model.Filter
	switch filter.Name {
	case "title":
		_filter = newTitleFilter(filter)
	case "fetch":
		_filter = newFetchFilter(filter)
	case "minify":
		_filter = newMinifyFilter(filter)
	default:
		// Try to load plugin regarding the name
		plug := plugin.GetRegsitry().LookupFilterPlugin(filter.Name)
		if plug == nil {
			return fmt.Errorf("unsuported filter: %s", filter.Name)
		}
		var err error
		_filter, err = plug.Build(filter.Props, filter.Tags)
		if err != nil {
			return fmt.Errorf("unable to create filter: %v", err)
		}
	}
	chain.filters = append(chain.filters, _filter)
	return nil
}

// Apply applies filter chain on an article
func (c *Chain) Apply(article *model.Article) error {
	for idx, filter := range c.filters {
		tags := filter.GetSpec().Tags
		if !article.Match(tags) {
			// Ignore filters that do not match the article tags
			continue
		}
		err := filter.DoFilter(article)
		if err != nil {
			return fmt.Errorf("error while applying filter #%d: %v", idx, err)
		}
	}
	return nil
}

// GetSpec return specification of the chain filter
func (c *Chain) GetSpec() []model.FilterSpec {
	result := make([]model.FilterSpec, len(c.filters))
	for idx, filter := range c.filters {
		result[idx] = filter.GetSpec()
	}
	return result
}
