package plugins

import "github.com/ncarlier/feedpushr/v3/pkg/model"

// GetBuiltinFilterPlugins get builtin plugins list
func GetBuiltinFilterPlugins() map[string]model.FilterPlugin {
	plugins := make(map[string]model.FilterPlugin)
	titleFilterPlugin := &TitleFilterPlugin{}
	minifyFilterPlugin := &MinifyFilterPlugin{}
	fetchFilterPlugin := &FetchFilterPlugin{}
	httpFilterPlugin := &HTTPFilterPlugin{}

	plugins[titleFilterPlugin.Spec().Name] = titleFilterPlugin
	plugins[minifyFilterPlugin.Spec().Name] = minifyFilterPlugin
	plugins[fetchFilterPlugin.Spec().Name] = fetchFilterPlugin
	plugins[httpFilterPlugin.Spec().Name] = httpFilterPlugin
	return plugins
}
