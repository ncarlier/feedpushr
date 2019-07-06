package plugin

import (
	"fmt"
	_plugin "plugin"

	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/rs/zerolog/log"
)

// Registry contains registered output and filter plugins
type Registry struct {
	outputPlugins map[string]model.OutputPlugin
	filterPlugins map[string]model.FilterPlugin
}

var instance *Registry

// GetRegsitry returns global plugin registry
func GetRegsitry() *Registry {
	if instance == nil {
		Configure([]string{})
	}
	return instance
}

// Configure global plugin registry
func Configure(plugins []string) error {
	reg := &Registry{
		outputPlugins: make(map[string]model.OutputPlugin),
		filterPlugins: make(map[string]model.FilterPlugin),
	}
	for _, filename := range plugins {
		plug, err := _plugin.Open(filename)
		if err != nil {
			return fmt.Errorf("unsuported plugin file: %s - %v", filename, err)
		}
		getPluginInfo, err := plug.Lookup("GetPluginInfo")
		if err != nil {
			return fmt.Errorf("unsuported plugin type: %s - %v", filename, err)
		}
		info := getPluginInfo.(func() model.PluginInfo)()
		log.Debug().Str("name", info.Name).Str("filename", filename).Msg("loading plugin...")

		switch info.Type {
		case model.OUTPUT_PLUGIN:
			getOutputPlugin, err := plug.Lookup("GetOutputPlugin")
			if err != nil {
				return fmt.Errorf("unsuported output plugin: %s - %v", info.Name, err)
			}
			outputPlugin, err := getOutputPlugin.(func() (model.OutputPlugin, error))()
			if err != nil {
				return fmt.Errorf("unable to load ouput plugin: %s - %v", info.Name, err)
			}
			reg.outputPlugins[info.Name] = outputPlugin
		case model.FILTER_PLUGIN:
			getFilter, err := plug.Lookup("GetFilterPlugin")
			if err != nil {
				return fmt.Errorf("unsuported filter plugin: %s - %v", info.Name, err)
			}
			filterPlugin, err := getFilter.(func() (model.FilterPlugin, error))()
			if err != nil {
				return fmt.Errorf("unable to load filter plugin: %s - %v", info.Name, err)
			}
			reg.filterPlugins[info.Name] = filterPlugin
		default:
			return fmt.Errorf("plugin type unknown: %d", info.Type)
		}
	}

	instance = reg
	return nil
}

// LookupOutputPlugin retrieve an output plugin by its name
func (r *Registry) LookupOutputPlugin(name string) model.OutputPlugin {
	plug, ok := r.outputPlugins[name]
	if !ok {
		return nil
	}
	return plug
}

// LookupFilterPlugin retrieve a filter plugin by its name
func (r *Registry) LookupFilterPlugin(name string) model.FilterPlugin {
	plug, ok := r.filterPlugins[name]
	if !ok {
		return nil
	}
	return plug
}
