package output

import (
	"fmt"

	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/model"
)

// Add an output
func (m *Manager) Add(def *model.OutputDef) (model.OutputProvider, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	m.log.Debug().Str("name", def.Name).Msg("adding output...")

	plug, ok := m.plugins[def.Name]
	if !ok {
		return nil, fmt.Errorf("unsuported output provider: %s", def.Name)
	}

	nextID := 0
	for _, _provider := range m.providers {
		if nextID < _provider.GetDef().ID {
			nextID = _provider.GetDef().ID
		}
	}
	def.ID = nextID + 1
	provider, err := plug.Build(def)
	if err != nil {
		return nil, err
	}
	m.providers = append(m.providers, provider)
	m.log.Info().Str("name", def.Name).Msg("output added")
	return provider, nil
}

// Update an output
func (m *Manager) Update(output *model.OutputDef) (model.OutputProvider, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for idx, provider := range m.providers {
		if output.ID == provider.GetDef().ID {
			// TODO merge objects
			output.Name = provider.GetDef().Name
			m.log.Debug().Int("id", output.ID).Msg("updating output...")
			plug, ok := m.plugins[output.Name]
			if !ok {
				return nil, fmt.Errorf("unsuported output provider: %s", output.Name)
			}
			p, err := plug.Build(output)
			if err != nil {
				return nil, err
			}
			m.providers[idx] = p
			m.log.Info().Int("id", output.ID).Msg("output updated")
			return p, nil
		}
	}
	return nil, common.ErrOutputNotFound
}

// Remove an output
func (m *Manager) Remove(output *model.OutputDef) error {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for idx, provider := range m.providers {
		if output.ID == provider.GetDef().ID {
			m.log.Debug().Int("id", output.ID).Msg("removing output...")
			m.providers = append(m.providers[:idx], m.providers[idx+1:]...)
			m.log.Info().Int("id", output.ID).Msg("output removed")
			return nil
		}
	}
	return common.ErrOutputNotFound
}

// Get an output
func (m *Manager) Get(id int) (model.OutputProvider, error) {
	for _, provider := range m.providers {
		if id == provider.GetDef().ID {
			return provider, nil
		}
	}
	return nil, common.ErrOutputNotFound
}
