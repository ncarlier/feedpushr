package format

import (
	"fmt"

	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

// NewOutputFormatter create formatter from output plugin properties
func NewOutputFormatter(output *model.OutputDef) (Formatter, error) {
	var formatter Formatter
	var formatValue string
	if formatProp, ok := output.Props["format"]; ok && formatProp != "" {
		formatValue = fmt.Sprintf("%v", formatProp)
		var err error
		formatter, err = NewTemplateFormatter(output.Hash(), formatValue)
		if err != nil {
			return nil, err
		}
	} else {
		formatter = NewJSONFormatter()
	}
	return formatter, nil
}
