package model

// PropType is a enum to specify a property type
type PropType int

const (
	// Email type
	Email PropType = iota
	// Number type
	Number
	// Password type
	Password
	// Text type
	Text
	// URL type
	URL
	// Select type
	Select
	// Textarea type
	Textarea
)

func (p PropType) String() string {
	return [...]string{"email", "number", "password", "text", "url", "select", "textarea"}[p]
}

// Spec describe specifications of a processor
type Spec struct {
	Name      string     `json:"name"`
	Desc      string     `json:"-"`
	PropsSpec []PropSpec `json:"-"`
}

// PropSpec contains property specification
type PropSpec struct {
	Desc    string
	Name    string
	Type    PropType
	Options map[string]string
}
