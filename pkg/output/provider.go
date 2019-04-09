package output

import (
	"fmt"
	"net/url"

	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/feedpushr/pkg/plugin"
	"github.com/rs/zerolog/log"
)

// newOutputProvider creates new output provider.
func newOutputProvider(uri string, pr *plugin.Registry) (model.OutputProvider, error) {
	logger := log.With().Str("component", "output").Logger()
	if uri == "" {
		uri = "stdout://"
	}
	u, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("invalid output URL: %s", uri)
	}
	var provider model.OutputProvider
	switch u.Scheme {
	case "stdout":
		provider = newStdOutputProvider(u.Fragment)
		logger.Info().Msg("using STDOUT output provider")
	case "http", "https":
		provider = newHTTPOutputProvider(uri, u.Fragment)
		logger.Info().Str("url", uri).Msg("using HTTP output provider")
	default:
		// Try to load plugin regarding the scheme
		plug := pr.LookupOutputPlugin(u.Scheme)
		if plug == nil {
			return nil, fmt.Errorf("unsuported output provider: %s", u.Scheme)
		}
		provider, err = plug.Build(u.Query())
		if err != nil {
			return nil, fmt.Errorf("unable to create output provider: %v", err)
		}
		logger.Info().Str("url", uri).Str("provider", u.Scheme).Msg("using external output provider")
	}
	return provider, nil
}
