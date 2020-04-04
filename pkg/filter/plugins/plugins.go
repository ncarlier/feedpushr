package plugins

import "github.com/ncarlier/feedpushr/v2/pkg/model"

// GetBuiltinFilterPlugins get builtin plugins list
func GetBuiltinFilterPlugins() map[string]model.FilterPlugin {
	plugins := make(map[string]model.FilterPlugin)
	titleFilterPlugin := &TitleFilterPlugin{}
	minifyFilterPlugin := &MinifyFilterPlugin{}
	fetchFilterPlugin := &FetchFilterPlugin{}

	plugins[titleFilterPlugin.Spec().Name] = titleFilterPlugin
	plugins[minifyFilterPlugin.Spec().Name] = minifyFilterPlugin
	plugins[fetchFilterPlugin.Spec().Name] = fetchFilterPlugin
	return plugins
}
