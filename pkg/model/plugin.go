package model

// PluginType is the plugin type qualifier
type PluginType int

const (
	// OUTPUT_PLUGIN output plugin type
	OUTPUT_PLUGIN PluginType = iota
	// FILTER_PLUGIN filter plugin type
	FILTER_PLUGIN
)

// PluginInfo contains plugins informations
type PluginInfo struct {
	Name string
	Type PluginType
}
