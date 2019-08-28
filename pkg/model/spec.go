package model

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
	Type string
}
