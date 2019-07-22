package builder

import (
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/model"
)

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
