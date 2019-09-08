package builder

import (
	"net/url"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/model"
)

// FilterBuilder is a builder to create an Filter definition
type FilterBuilder struct {
	filter *model.FilterDef
}

// NewFilterBuilder creates new Filter definition builder instance
func NewFilterBuilder() *FilterBuilder {
	filter := &model.FilterDef{
		Props: make(model.FilterProps),
	}
	return &FilterBuilder{filter}
}

// Build creates the filter definition
func (fb *FilterBuilder) Build() *model.FilterDef {
	return fb.filter
}

// From creates filter form an other
func (fb *FilterBuilder) From(source model.FilterDef) *FilterBuilder {
	clone := source
	copy(clone.Tags, source.Tags)
	fb.filter = &clone
	return fb
}

// FromURI creates a filter definition form an URI
func (fb *FilterBuilder) FromURI(URI string) *FilterBuilder {
	u, err := url.Parse(URI)
	if err != nil {
		return fb
	}
	tags := GetFeedTags(&u.Fragment)
	for key, value := range u.Query() {
		fb.filter.Props[key] = value[0]
	}
	fb.filter.Name = u.Scheme
	fb.filter.Alias = u.Scheme
	fb.filter.Tags = tags
	fb.filter.Enabled = true
	return fb
}

// ID set ID
func (fb *FilterBuilder) ID(ID int) *FilterBuilder {
	fb.filter.ID = ID
	return fb
}

// Alias set alias
func (fb *FilterBuilder) Alias(alias *string) *FilterBuilder {
	if alias != nil {
		fb.filter.Alias = *alias
	}
	return fb
}

// Spec set spec name
func (fb *FilterBuilder) Spec(name string) *FilterBuilder {
	fb.filter.Name = name
	return fb
}

// Tags set tags
func (fb *FilterBuilder) Tags(tags *string) *FilterBuilder {
	if tags != nil {
		fb.filter.Tags = GetFeedTags(tags)
	}
	return fb
}

// Props set props
func (fb *FilterBuilder) Props(props model.FilterProps) *FilterBuilder {
	fb.filter.Props = props
	return fb
}

// Enable set enabled status
func (fb *FilterBuilder) Enable(status bool) *FilterBuilder {
	fb.filter.Enabled = status
	return fb
}

// NewFilterFromDef creates new Filter from a definition
func NewFilterFromDef(def model.FilterDef) *app.Filter {
	return &app.Filter{
		ID:      def.ID,
		Alias:   def.Alias,
		Name:    def.Name,
		Desc:    def.Desc,
		Props:   def.Props,
		Tags:    def.Tags,
		Enabled: def.Enabled,
	}
}
