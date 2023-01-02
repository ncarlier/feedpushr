package filter

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// Builder is used to create an Filter definition
type Builder struct {
	filter *model.FilterDef
}

// NewBuilder creates new Filter definition builder instance
func NewBuilder() *Builder {
	filter := &model.FilterDef{
		Props: make(model.FilterProps),
	}
	return &Builder{filter}
}

// Build creates the filter definition
func (fb *Builder) Build() *model.FilterDef {
	return fb.filter
}

// From creates filter form an other
func (fb *Builder) From(source model.FilterDef) *Builder {
	clone := source
	fb.filter = &clone
	return fb
}

// FromURI creates a filter definition form an URI
func (fb *Builder) FromURI(URI string) *Builder {
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
	return fb.NewID()
}

// ID set ID
func (fb *Builder) ID(ID string) *Builder {
	fb.filter.ID = ID
	return fb
}

// NewID set new ID
func (fb *Builder) NewID() *Builder {
	fb.filter.ID = uuid.New().String()
	return fb
}

// Alias set alias
func (fb *Builder) Alias(alias *string) *Builder {
	if alias != nil {
		fb.filter.Alias = *alias
	}
	return fb
}

// Spec set spec name
func (fb *Builder) Spec(name string) *Builder {
	fb.filter.Name = name
	return fb
}

// Condition set condition
func (fb *Builder) Condition(condition *string) *Builder {
	if condition != nil {
		fb.filter.Condition = *condition
	}
	return fb
}

// Props set props
func (fb *Builder) Props(props model.FilterProps) *Builder {
	if len(props) > 0 {
		fb.filter.Props = props
	}
	return fb
}

// Enable set enabled status
func (fb *Builder) Enable(status bool) *Builder {
	fb.filter.Enabled = status
	return fb
}

// NewFilterResponseFromDef creates new Filter response from a definition
func NewFilterResponseFromDef(def *model.FilterDef) *app.FilterResponse {
	if def == nil {
		return nil
	}
	return &app.FilterResponse{
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
}
