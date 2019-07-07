package builder

import (
	"fmt"
	"net/url"

	"github.com/ncarlier/feedpushr/autogen/app"
)

// NewFilterFromURI create new filter from URI definition
func NewFilterFromURI(URI string) (*app.Filter, error) {
	u, err := url.Parse(URI)
	if err != nil {
		return nil, fmt.Errorf("invalid filter URI: %s", URI)
	}
	tags := GetFeedTags(&u.Fragment)
	props := make(map[string]interface{})
	for key, value := range u.Query() {
		props[key] = value[0]
	}
	return &app.Filter{
		Name:  u.Scheme,
		Props: props,
		Tags:  tags,
	}, nil
}
