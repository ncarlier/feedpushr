package pipeline

import (
	"fmt"

	"github.com/ncarlier/feedpushr/v2/pkg/common"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

// AddOutput add an output to the pipeline
func (p *Pipeline) AddOutput(def *model.OutputDef) (model.Output, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	p.log.Debug().Str("name", def.Name).Msg("adding output...")

	plug, ok := p.plugins[def.Name]
	if !ok {
		return nil, fmt.Errorf("unsupported output provider: %s", def.Name)
	}

	nextID := 0
	for _, _provider := range p.outputs {
		if nextID < _provider.GetDef().ID {
			nextID = _provider.GetDef().ID
		}
	}
	def.ID = nextID + 1
	output, err := plug.Build(def)
	if err != nil {
		return nil, err
	}
	p.outputs = append(p.outputs, output)
	p.log.Info().Str("name", def.Name).Msg("output added")
	return output, nil
}

// UpdateOutput update an output of the pipeline
func (p *Pipeline) UpdateOutput(output *model.OutputDef) (model.Output, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	for idx, provider := range p.outputs {
		if output.ID == provider.GetDef().ID {
			// TODO merge objects
			output.Name = provider.GetDef().Name
			p.log.Debug().Int("id", output.ID).Msg("updating output...")
			plug, ok := p.plugins[output.Name]
			if !ok {
				return nil, fmt.Errorf("unsupported output provider: %s", output.Name)
			}
			provider, err := plug.Build(output)
			if err != nil {
				return nil, err
			}
			p.outputs[idx] = provider
			p.log.Info().Int("id", output.ID).Msg("output updated")
			return provider, nil
		}
	}
	return nil, common.ErrOutputNotFound
}

// RemoveOutput removes an output from the pipeline
func (p *Pipeline) RemoveOutput(output *model.OutputDef) error {
	p.lock.RLock()
	defer p.lock.RUnlock()

	for idx, provider := range p.outputs {
		if output.ID == provider.GetDef().ID {
			p.log.Debug().Int("id", output.ID).Msg("removing output...")
			p.outputs = append(p.outputs[:idx], p.outputs[idx+1:]...)
			p.log.Info().Int("id", output.ID).Msg("output removed")
			return nil
		}
	}
	return common.ErrOutputNotFound
}

// GetOutput retrieve an output of the pipeline
func (p *Pipeline) GetOutput(id int) (model.Output, error) {
	for _, provider := range p.outputs {
		if id == provider.GetDef().ID {
			return provider, nil
		}
	}
	return nil, common.ErrOutputNotFound
}

// GetOutputDefs return all output definitions of the pipeline
func (p *Pipeline) GetOutputDefs() []model.OutputDef {
	result := make([]model.OutputDef, len(p.outputs))
	for idx, provider := range p.outputs {
		result[idx] = provider.GetDef()
	}
	return result
}
