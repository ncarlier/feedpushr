package service

import (
	"fmt"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/filter"
	"github.com/ncarlier/feedpushr/pkg/store"
)

func loadChainFilter(db store.DB) (*filter.Chain, error) {
	chain := filter.NewChainFilter()
	err := db.ForEachFilter(func(f *app.Filter) error {
		if f == nil {
			return fmt.Errorf("filter is null")
		}
		chain.Add(f)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return chain, nil
}
