package model

type PropType int

const (
	Email PropType = iota
	Number
	Password
	Text
	URL
)

func (p PropType) String() string {
	return [...]string{"email", "number", "password", "text", "url"}[p]
}

// Spec describe specifications of a processor
type Spec struct {
	Name      string     `json:"name"`
	Desc      string     `json:"-"`
	PropsSpec []PropSpec `json:"-"`
}

// PropSpec contains property specification
type PropSpec struct {
	Desc string
	Name string
	Type PropType
}
