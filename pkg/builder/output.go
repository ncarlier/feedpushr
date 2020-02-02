package builder

import (
	"net/url"

	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
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
	ob.output = &clone
	return ob
}

// FromURI creates a output definition form an URI
func (ob *OutputBuilder) FromURI(URI string) *OutputBuilder {
	u, err := url.Parse(URI)
	if err != nil {
		return ob
	}
	for key, value := range u.Query() {
		ob.output.Props[key] = value[0]
	}
	ob.output.Name = u.Scheme
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

// Condition set condition
func (ob *OutputBuilder) Condition(condition *string) *OutputBuilder {
	if condition != nil {
		ob.output.Condition = *condition
	}
	return ob
}

// Props set props
func (ob *OutputBuilder) Props(props model.OutputProps) *OutputBuilder {
	if len(props) > 0 {
		ob.output.Props = props
	}
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
		ID:        def.ID,
		Alias:     def.Alias,
		Name:      def.Name,
		Desc:      def.Desc,
		Props:     def.Props,
		Condition: def.Condition,
		Enabled:   def.Enabled,
	}
}
