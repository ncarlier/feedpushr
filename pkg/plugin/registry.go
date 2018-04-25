package plugin

import (
	"fmt"
	_plugin "plugin"

	"github.com/ncarlier/feedpushr/pkg/model"
)

// OutputPlugin is the scructure of an ouptut plugin
type OutputPlugin struct {
	Name     string
	Provider model.OutputProvider
}

// FilterPlugin is the scructure of an filter plugin
type FilterPlugin struct {
	Name   string
	Filter model.Filter
}

// Registry contains registered output and filter plugins
type Registry struct {
	outputPlugins map[string]OutputPlugin
	filterPlugins map[string]FilterPlugin
}

// NewPluginRegistry creates a new plugin registry
func NewPluginRegistry(plugins []string) (*Registry, error) {
	reg := &Registry{
		outputPlugins: make(map[string]OutputPlugin),
		filterPlugins: make(map[string]FilterPlugin),
	}
	for _, filename := range plugins {
		plug, err := _plugin.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("unsuported plugin file: %s - %v", filename, err)
		}
		getPluginInfo, err := plug.Lookup("GetPluginInfo")
		if err != nil {
			return nil, fmt.Errorf("unsuported plugin type: %s - %v", filename, err)
		}
		info := getPluginInfo.(func() model.PluginInfo)()

		switch info.Type {
		case model.OUTPUT_PLUGIN:
			getOutputProvider, err := plug.Lookup("GetOutputProvider")
			if err != nil {
				return nil, fmt.Errorf("unsuported output plugin: %s - %v", info.Name, err)
			}
			provider, err := getOutputProvider.(func() (model.OutputProvider, error))()
			if err != nil {
				return nil, fmt.Errorf("unable to configure pugin output provider: %s - %v", info.Name, err)
			}
			reg.outputPlugins[info.Name] = OutputPlugin{
				Name:     info.Name,
				Provider: provider,
			}
		case model.FILTER_PLUGIN:
			getFilter, err := plug.Lookup("GetFilter")
			if err != nil {
				return nil, fmt.Errorf("unsuported filter plugin: %s - %v", info.Name, err)
			}
			filter, err := getFilter.(func() (model.Filter, error))()
			if err != nil {
				return nil, fmt.Errorf("unable to configure pugin filter: %s - %v", info.Name, err)
			}
			reg.filterPlugins[info.Name] = FilterPlugin{
				Name:   info.Name,
				Filter: filter,
			}
		default:
			return nil, fmt.Errorf("plugin type unknown: %d", info.Type)
		}
	}

	return reg, nil
}

// LookupOutputPlugin retrieve an output plugin by its name
func (r *Registry) LookupOutputPlugin(name string) *OutputPlugin {
	plug, ok := r.outputPlugins[name]
	if !ok {
		return nil
	}
	return &plug
}

// LookupFilterPlugin retrieve a filter plugin by its name
func (r *Registry) LookupFilterPlugin(name string) *FilterPlugin {
	plug, ok := r.filterPlugins[name]
	if !ok {
		return nil
	}
	return &plug
}
