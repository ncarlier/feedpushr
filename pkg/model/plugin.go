package model

// PluginType is the plugin type qualifier
type PluginType int

const (
	// OutputPluginType output plugin type
	OutputPluginType PluginType = iota
	// FilterPluginType filter plugin type
	FilterPluginType
)

// PluginSpec contains plugins specifications
type PluginSpec struct {
	Type PluginType
	Spec
}

// OutputPlugin is the interface of an output plugin
type OutputPlugin interface {
	// Build an output plugin
	Build(def *OutputDef) (Output, error)
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
