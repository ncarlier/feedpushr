package server

import (
	"context"
	"fmt"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/output"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
)

func loadOutputs(db store.DB, om *output.Manager) error {
	// Load output outputs from DB
	return db.ForEachOutput(context.Background(), func(o *model.OutputDef) error {
		if o == nil {
			return fmt.Errorf("output is null")
		}
		_, err := om.AddOutputProcessor(o)
		return err
	})
}
