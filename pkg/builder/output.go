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
func (fb *OutputBuilder) Build() *model.OutputDef {
	return fb.output
}

// FromURI creates a output definition form an URI
func (fb *OutputBuilder) FromURI(URI string) *OutputBuilder {
	u, err := url.Parse(URI)
	if err != nil {
		return fb
	}
	tags := GetFeedTags(&u.Fragment)
	for key, value := range u.Query() {
		fb.output.Props[key] = value[0]
	}
	fb.output.Name = u.Scheme
	fb.output.Tags = tags
	fb.output.Enabled = true
	return fb
}

// ID set ID
func (fb *OutputBuilder) ID(ID int) *OutputBuilder {
	fb.output.ID = ID
	return fb
}

// Spec set spec name
func (fb *OutputBuilder) Spec(name string) *OutputBuilder {
	fb.output.Name = name
	return fb
}

// Tags set tags
func (fb *OutputBuilder) Tags(tags *string) *OutputBuilder {
	fb.output.Tags = GetFeedTags(tags)
	return fb
}

// Props set props
func (fb *OutputBuilder) Props(props model.OutputProps) *OutputBuilder {
	fb.output.Props = props
	return fb
}

// Enable set enabled status
func (fb *OutputBuilder) Enable(status bool) *OutputBuilder {
	fb.output.Enabled = status
	return fb
}

// NewOutputFromDef creates new Output from a definition
func NewOutputFromDef(def model.OutputDef) *app.Output {
	return &app.Output{
		ID:      def.ID,
		Name:    def.Name,
		Desc:    def.Desc,
		Props:   def.Props,
		Tags:    def.Tags,
		Enabled: def.Enabled,
	}
}
