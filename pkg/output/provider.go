package output

import (
	"fmt"
	"net/url"

	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/model"
	"github.com/ncarlier/feedpushr/pkg/plugin"
	"github.com/rs/zerolog/log"
)

// newOutputProvider creates new output provider.
func newOutputProvider(uri string) (model.OutputProvider, error) {
	logger := log.With().Str("component", "output").Logger()
	if uri == "" {
		uri = "stdout://"
	}
	u, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("invalid output URL: %s", uri)
	}
	tags := builder.GetFeedTags(&u.Fragment)
	var provider model.OutputProvider
	switch u.Scheme {
	case "stdout":
		provider = newStdOutputProvider(tags)
		logger.Info().Msg("using STDOUT output provider")
	case "http", "https":
		provider = newHTTPOutputProvider(uri, tags)
		logger.Info().Str("url", uri).Msg("using HTTP output provider")
	default:
		// Try to load plugin regarding the scheme
		plug := plugin.GetRegsitry().LookupOutputPlugin(u.Scheme)
		if plug == nil {
			return nil, fmt.Errorf("unsuported output provider: %s", u.Scheme)
		}
		provider, err = plug.Build(u.Query(), tags)
		if err != nil {
			return nil, fmt.Errorf("unable to create output provider: %v", err)
		}
		logger.Info().Str("url", uri).Str("provider", u.Scheme).Msg("using external output provider")
	}
	return provider, nil
}
