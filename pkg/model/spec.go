package model

type Spec struct {
	Name      string
	Desc      string
	PropsSpec []PropSpec
}

// PropSpec contains property specification
type PropSpec struct {
	Desc string
	Name string
	Type string
}
