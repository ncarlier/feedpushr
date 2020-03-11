package fn

import (
	"strings"
	"text/template"
)

// Functions used inside templates
var Functions = template.FuncMap{
	"upper":    strings.ToUpper,
	"tweet":    Tweet,
	"truncate": Truncate,
}
