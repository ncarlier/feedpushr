package model

// PluginType is the plugin type qualifier
type PluginType int

const (
	// OUTPUT_PLUGIN output plugin type
	OUTPUT_PLUGIN PluginType = iota
	// FILTER_PLUGIN filter plugin type
	FILTER_PLUGIN
)

// PluginSpec contains plugins specifications
type PluginSpec struct {
	Type PluginType
	Spec
}

// OutputPlugin is the interface of an output plugin
type OutputPlugin interface {
	// Build an output plugin
	Build(def *OutputDef) (OutputProvider, error)
	// Spec returns plugin specs
	Spec() Spec
}

// FilterPlugin is the interface of an filter plugin
type FilterPlugin interface {
	// Build a filter
	Build(def *FilterDef) (Filter, error)
	// Spec returns plugin specs
	Spec() Spec
}
