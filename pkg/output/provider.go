package output

import (
	"fmt"

	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/feedpushr/pkg/plugin"
)

// newOutputProvider creates new output provider.
func newOutputProvider(output *app.Output) (model.OutputProvider, error) {
	var provider model.OutputProvider
	switch output.Name {
	case "stdout":
		provider = newStdOutputProvider(output)
	case "http":
		provider = newHTTPOutputProvider(output)
	default:
		// Try to load plugin regarding the scheme
		plug := plugin.GetRegsitry().LookupOutputPlugin(output.Name)
		if plug == nil {
			return nil, fmt.Errorf("unsuported output provider: %s", output.Name)
		}
		var err error
		provider, err = plug.Build(output.Props, output.Tags)
		if err != nil {
			return nil, fmt.Errorf("unable to create output provider: %v", err)
		}
	}
	return provider, nil
}

// Add an output
func (m *Manager) Add(output *app.Output) error {
	m.lock.RLock()
	defer m.lock.RUnlock()
	m.log.Debug().Str("name", output.Name).Msg("adding output...")
	provider, err := newOutputProvider(output)
	if err != nil {
		return err
	}
	m.providers = append(m.providers, provider)
	m.log.Debug().Str("name", output.Name).Msg("output added")
	return nil
}

// Update an output
func (m *Manager) Update(output *app.Output) error {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for idx, provider := range m.providers {
		if output.ID == provider.GetDef().ID {
			p, err := newOutputProvider(output)
			if err != nil {
				return err
			}
			m.providers[idx] = p
			return nil
		}
	}
	return common.ErrOutputNotFound
}

// Remove an output
func (m *Manager) Remove(output *app.Output) error {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for idx, provider := range m.providers {
		if output.ID == provider.GetDef().ID {
			m.providers = append(m.providers[:idx], m.providers[idx+1:]...)
			return nil
		}
	}
	return common.ErrOutputNotFound
}
