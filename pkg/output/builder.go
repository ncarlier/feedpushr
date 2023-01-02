package output

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
	"github.com/ncarlier/feedpushr/v3/pkg/filter"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// Builder is used to create an Output definition
type Builder struct {
	output *model.OutputDef
}

// NewBuilder creates new Output definition builder instance
func NewBuilder() *Builder {
	output := &model.OutputDef{
		Props: make(model.OutputProps),
	}
	return &Builder{output}
}

// Build creates the output definition
func (ob *Builder) Build() *model.OutputDef {
	return ob.output
}

// From creates output form an other
func (ob *Builder) From(source model.OutputDef) *Builder {
	clone := source
	ob.output = &clone
	return ob
}

// FromURI creates a output definition form an URI
func (ob *Builder) FromURI(URI string) *Builder {
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
func (ob *Builder) ID(ID string) *Builder {
	ob.output.ID = ID
	return ob
}

// NewID set new ID
func (ob *Builder) NewID() *Builder {
	ob.output.ID = uuid.New().String()
	return ob
}

// Alias set alias
func (ob *Builder) Alias(alias *string) *Builder {
	if alias != nil {
		ob.output.Alias = *alias
	}
	return ob
}

// Spec set spec name
func (ob *Builder) Spec(name string) *Builder {
	ob.output.Name = name
	return ob
}

// Condition set condition
func (ob *Builder) Condition(condition *string) *Builder {
	if condition != nil {
		ob.output.Condition = *condition
	}
	return ob
}

// Props set props
func (ob *Builder) Props(props model.OutputProps) *Builder {
	if len(props) > 0 {
		ob.output.Props = props
	}
	return ob
}

// Enable set enabled status
func (ob *Builder) Enable(status bool) *Builder {
	ob.output.Enabled = status
	return ob
}

// NewOutputResponseFromDef creates new Output response from a definition
func NewOutputResponseFromDef(def *model.OutputDef) *app.OutputResponse {
	if def == nil {
		return nil
	}
	result := app.OutputResponse{
		ID:        def.ID,
		Alias:     def.Alias,
		Name:      def.Name,
		Desc:      def.Desc,
		Props:     def.Props,
		Condition: def.Condition,
		Enabled:   def.Enabled,
		NbSuccess: int(def.NbSuccess),
		NbError:   int(def.NbError),
	}

	for _, filterDef := range def.Filters {
		result.Filters = append(result.Filters, filter.NewFilterResponseFromDef(filterDef))
	}

	return &result
}
