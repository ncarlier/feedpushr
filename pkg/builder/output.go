package builder

import (
	"net/url"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/model"
)

// OutputBuilder is a builder to create an Output definition
type OutputBuilder struct {
	output *model.OutputDef
}

// NewOutputBuilder creates new Output definition builder instance
func NewOutputBuilder() *OutputBuilder {
	output := &model.OutputDef{
		Props: make(model.OutputProps),
	}
	return &OutputBuilder{output}
}

// Build creates the output definition
func (ob *OutputBuilder) Build() *model.OutputDef {
	return ob.output
}

// From creates output form an other
func (ob *OutputBuilder) From(source model.OutputDef) *OutputBuilder {
	clone := source
	copy(clone.Tags, source.Tags)
	ob.output = &clone
	return ob
}

// FromURI creates a output definition form an URI
func (ob *OutputBuilder) FromURI(URI string) *OutputBuilder {
	u, err := url.Parse(URI)
	if err != nil {
		return ob
	}
	tags := GetFeedTags(&u.Fragment)
	for key, value := range u.Query() {
		ob.output.Props[key] = value[0]
	}
	ob.output.Name = u.Scheme
	ob.output.Tags = tags
	ob.output.Enabled = true
	return ob
}

// ID set ID
func (ob *OutputBuilder) ID(ID int) *OutputBuilder {
	ob.output.ID = ID
	return ob
}

// Alias set alias
func (ob *OutputBuilder) Alias(alias *string) *OutputBuilder {
	if alias != nil {
		ob.output.Alias = *alias
	}
	return ob
}

// Spec set spec name
func (ob *OutputBuilder) Spec(name string) *OutputBuilder {
	ob.output.Name = name
	return ob
}

// Tags set tags
func (ob *OutputBuilder) Tags(tags *string) *OutputBuilder {
	if tags != nil {
		ob.output.Tags = GetFeedTags(tags)
	}
	return ob
}

// Props set props
func (ob *OutputBuilder) Props(props model.OutputProps) *OutputBuilder {
	ob.output.Props = props
	return ob
}

// Enable set enabled status
func (ob *OutputBuilder) Enable(status bool) *OutputBuilder {
	ob.output.Enabled = status
	return ob
}

// NewOutputFromDef creates new Output from a definition
func NewOutputFromDef(def model.OutputDef) *app.Output {
	return &app.Output{
		ID:      def.ID,
		Alias:   def.Alias,
		Name:    def.Name,
		Desc:    def.Desc,
		Props:   def.Props,
		Tags:    def.Tags,
		Enabled: def.Enabled,
	}
}
