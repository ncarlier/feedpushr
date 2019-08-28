package service

import (
	"fmt"

	"github.com/ncarlier/feedpushr/pkg/filter"
	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/feedpushr/pkg/store"
)

func loadChainFilter(db store.DB) (*filter.Chain, error) {
	chain := filter.NewChainFilter()
	err := db.ForEachFilter(func(f *model.FilterDef) error {
		if f == nil {
			return fmt.Errorf("filter is null")
		}
		_, err := chain.Add(f)
		return err
	})
	if err != nil {
		return nil, err
	}
	return chain, nil
}
