package builder

import (
	"net/url"

	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
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
	fb.filter = &clone
	return fb
}

// FromURI creates a filter definition form an URI
func (fb *FilterBuilder) FromURI(URI string) *FilterBuilder {
	u, err := url.Parse(URI)
	if err != nil {
		return fb
	}
	for key, value := range u.Query() {
		fb.filter.Props[key] = value[0]
	}
	fb.filter.Name = u.Scheme
	fb.filter.Alias = u.Scheme
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

// Condition set condition
func (fb *FilterBuilder) Condition(condition *string) *FilterBuilder {
	if condition != nil {
		fb.filter.Condition = *condition
	}
	return fb
}

// Props set props
func (fb *FilterBuilder) Props(props model.FilterProps) *FilterBuilder {
	if len(props) > 0 {
		fb.filter.Props = props
	}
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
		ID:        def.ID,
		Alias:     def.Alias,
		Name:      def.Name,
		Desc:      def.Desc,
		Props:     def.Props,
		Condition: def.Condition,
		Enabled:   def.Enabled,
	}
}
