package builder

import (
	"fmt"
	"net/url"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/model"
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
		Name:    u.Scheme,
		Props:   props,
		Tags:    tags,
		Enabled: true,
	}, nil
}

// NewFilterFromDef creates new Filter from a definition
func NewFilterFromDef(def model.FilterDef) *app.Filter {
	return &app.Filter{
		ID:      def.ID,
		Name:    def.Name,
		Desc:    def.Desc,
		Props:   def.Props,
		Tags:    def.Tags,
		Enabled: def.Enabled,
	}
}
