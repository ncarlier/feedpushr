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

// Registry contains registered output and filter plugins
type Registry struct {
	outputPlugins map[string]OutputPlugin
}

// LocalRegistry is the local plugin registry
var LocalRegistry = &Registry{}

// LookupOutputPlugin retrieve an output plugin by its name
func LookupOutputPlugin(name string) *OutputPlugin {
	plug, ok := LocalRegistry.outputPlugins[name]
	if !ok {
		return nil
	}
	return &plug
}

// String get the registry string value
// This is used to match the flag interface
func (r *Registry) String() string {
	return "plugin registry"
}

// Set loads a plugin by its filename into the registry
// This is used to match the flag interface
func (r *Registry) Set(filename string) error {
	plug, err := _plugin.Open(filename)
	if err != nil {
		return fmt.Errorf("unsuported plugin file: %s - %v", filename, err)
	}
	getPluginInfo, err := plug.Lookup("GetPluginInfo")
	if err != nil {
		return fmt.Errorf("unsuported plugin type: %s - %v", filename, err)
	}
	info := getPluginInfo.(func() model.PluginInfo)()
	if info.Type == model.OUTPUT_PLUGIN {
		getOutputProvider, err := plug.Lookup("GetOutputProvider")
		if err != nil {
			return fmt.Errorf("unsuported output plugin: %s - %v", info.Name, err)
		}
		provider, err := getOutputProvider.(func() (model.OutputProvider, error))()
		if err != nil {
			return fmt.Errorf("unable to configure pugin output provider: %s - %v", info.Name, err)
		}
		r.outputPlugins[info.Name] = OutputPlugin{
			Name:     info.Name,
			Provider: provider,
		}
	}

	return nil
}
