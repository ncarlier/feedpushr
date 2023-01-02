package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/filter"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/output"
)

func TestNewManager(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	// New output definition
	def := output.NewBuilder().FromURI("stdout://").NewID().Enable(true).Build()
	// New filter
	f1 := filter.NewBuilder().FromURI("title://?prefix=hello").NewID().Build()
	// New chain filter
	chain, err := filter.NewChainFilter(model.FilterDefCollection{})
	assert.Nil(t, err)
	chain.Add(f1)
	// Add filters to output definition
	def.Filters = chain.GetFilterDefs()

	// Create new output manager
	manager, err := output.NewOutputManager(cm)
	assert.Nil(t, err)
	// Add output definition to the manager
	processor, err := manager.AddOutputProcessor(def)
	assert.Nil(t, err)
	assert.Equal(t, uint32(0), processor.GetDef().NbSuccess)
	assert.Equal(t, uint32(0), processor.GetDef().NbError)

	// Send articles to the manager
	now := time.Now()
	article := &model.Article{
		Title:           "World",
		PublishedParsed: &now,
		Tags:            []string{"test"},
	}
	manager.Push([]*model.Article{article})
	time.Sleep(100 * time.Millisecond)
	manager.Shutdown()
	assert.Equal(t, uint32(1), processor.GetDef().NbSuccess)
	assert.Equal(t, uint32(0), processor.GetDef().NbError)
}
