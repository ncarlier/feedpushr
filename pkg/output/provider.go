package output

import (
	"fmt"
	"net/url"
	"plugin"

	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/rs/zerolog/log"
)

// newOutputProvider creates new output provider.
func newOutputProvider(uri string) (model.OutputProvider, error) {
	logger := log.With().Str("component", "output").Logger()
	var scheme string
	if uri == "" || uri == "stdout" {
		scheme = "stdout"
	} else {
		u, err := url.ParseRequestURI(uri)
		if err != nil {
			return nil, fmt.Errorf("invalid output URL: %s", uri)
		}
		scheme = u.Scheme
	}
	var provider model.OutputProvider
	switch scheme {
	case "stdout":
		provider = newStdOutputProvider()
		logger.Info().Msg("using STDOUT output provider")
	case "http", "https":
		provider = newHTTPOutputProvider(uri)
		logger.Info().Str("url", uri).Msg("using HTTP output provider")
	default:
		// Try to load plugin regarding the scheme
		pluginName := fmt.Sprintf("feedpushr-%s.so", scheme)
		plug, err := plugin.Open(pluginName)
		if err != nil {
			return nil, fmt.Errorf("unsuported output provider: %s - %v", scheme, err)
		}
		getOutputProvider, err := plug.Lookup("GetOutputProvider")
		if err != nil {
			return nil, fmt.Errorf("unsuported output provider: %s - %v", scheme, err)
		}
		provider, err = getOutputProvider.(func() (model.OutputProvider, error))()
		if err != nil {
			return nil, fmt.Errorf("unsuported output provider: %s - %v", scheme, err)
		}
		logger.Info().Str("url", uri).Str("provider", scheme).Msg("using external output provider")
	}
	return provider, nil
}
