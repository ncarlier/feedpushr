package output

import "github.com/ncarlier/feedpushr/v2/pkg/model"

// GetBuiltinOutputPlugins get builtin plugins list
func GetBuiltinOutputPlugins() map[string]model.OutputPlugin {
	plugins := make(map[string]model.OutputPlugin)
	stdoutOutputPlugin := &StdoutOutputPlugin{}
	httpOutputPlugin := &HTTPOutputPlugin{}
	readflowOutputPlugin := &ReadflowOutputPlugin{}

	plugins[stdoutOutputPlugin.Spec().Name] = stdoutOutputPlugin
	plugins[httpOutputPlugin.Spec().Name] = httpOutputPlugin
	plugins[readflowOutputPlugin.Spec().Name] = readflowOutputPlugin
	return plugins
}
