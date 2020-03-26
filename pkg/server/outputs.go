package server

import (
	"fmt"

	"github.com/ncarlier/feedpushr/v2/pkg/model"
	"github.com/ncarlier/feedpushr/v2/pkg/output"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
)

func loadOutputs(db store.DB, om *output.Manager) error {
	// Load output outputs from DB
	return db.ForEachOutput(func(o *model.OutputDef) error {
		if o == nil {
			return fmt.Errorf("output is null")
		}
		_, err := om.AddOutputProcessor(o)
		return err
	})
}
