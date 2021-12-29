package plugins

import "github.com/ncarlier/feedpushr/v3/pkg/model"

// GetBuiltinOutputPlugins get builtin plugins list
func GetBuiltinOutputPlugins() map[string]model.OutputPlugin {
	plugins := make(map[string]model.OutputPlugin)
	stdoutOutputPlugin := &StdoutOutputPlugin{}
	httpOutputPlugin := &HTTPOutputPlugin{}
	emailOutputPlugin := &EmailOutputPlugin{}
	readflowOutputPlugin := &ReadflowOutputPlugin{}

	plugins[stdoutOutputPlugin.Spec().Name] = stdoutOutputPlugin
	plugins[httpOutputPlugin.Spec().Name] = httpOutputPlugin
	plugins[emailOutputPlugin.Spec().Name] = emailOutputPlugin
	plugins[readflowOutputPlugin.Spec().Name] = readflowOutputPlugin
	return plugins
}
