package output

import (
	"fmt"

	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/filter"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

// AddOutputProcessor add an output processor to the manager
func (m *Manager) AddOutputProcessor(def *model.OutputDef) (*Processor, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	m.logger.Debug().Str("name", def.Name).Msg("adding output processor...")

	plug, ok := m.plugins[def.Name]
	if !ok {
		return nil, fmt.Errorf("unsupported output provider: %s", def.Name)
	}

	out, err := plug.Build(def)
	if err != nil {
		return nil, err
	}

	chain, err := filter.NewChainFilter(def.Filters)
	if err != nil {
		return nil, err
	}

	processor, err := NewOutputProcessor(out, chain, m.cache)
	if err != nil {
		return nil, err
	}

	m.processors[def.ID] = processor
	m.logger.Info().Str("name", def.Name).Str("id", def.ID).Msg("output processor added")
	return processor, nil
}

// UpdateOutputProcessor update an output of the pipeline
func (m *Manager) UpdateOutputProcessor(def *model.OutputDef) (*Processor, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	m.logger.Debug().Str("name", def.Name).Msg("updating output processor...")

	processor, err := m.GetOutputProcessor(def.ID)
	if err != nil {
		return nil, err
	}

	plug, ok := m.plugins[def.Name]
	if !ok {
		return nil, fmt.Errorf("unsupported output provider: %s", def.Name)
	}

	out, err := plug.Build(def)
	if err != nil {
		return nil, err
	}

	chain, err := filter.NewChainFilter(def.Filters)
	if err != nil {
		return nil, err
	}

	processor.Update(out, chain)
	m.logger.Info().Str("name", def.Name).Str("id", def.ID).Msg("output processor updated")
	return processor, nil
}

// RemoveOutputProcessor removes an output from the pipeline
func (m *Manager) RemoveOutputProcessor(def *model.OutputDef) error {
	m.lock.RLock()
	defer m.lock.RUnlock()

	processor, err := m.GetOutputProcessor(def.ID)
	if err != nil {
		return err
	}
	m.logger.Debug().Str("id", def.ID).Msg("removing output procesor...")
	processor.Shutdown()
	delete(m.processors, def.ID)
	m.logger.Info().Str("id", def.ID).Msg("output processor removed")
	return nil
}

// GetOutputProcessor removes an output from the pipeline
func (m *Manager) GetOutputProcessor(id string) (*Processor, error) {
	processor, found := m.processors[id]
	if !found {
		return nil, common.ErrOutputNotFound
	}
	return processor, nil
}

// GetOutputDefs return all output definitions of the pipeline
func (m *Manager) GetOutputDefs() []model.OutputDef {
	result := []model.OutputDef{}
	for _, processor := range m.processors {
		result = append(result, processor.GetDef())
	}
	return result
}
