package output

import (
	"fmt"
	"net/url"

	"github.com/mmcdole/gofeed"
	"github.com/rs/zerolog/log"
)

// Provider is the output provider interface
type Provider interface {
	Send(article *gofeed.Item) error
}

// newOutputProvider creates new output provider.
func newOutputProvider(uri string) (Provider, error) {
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
	var provider Provider
	switch scheme {
	case "stdout":
		provider = newStdOutputProvider()
		logger.Info().Msg("using STDOUT output provider")
	case "http", "https":
		provider = newHTTPOutputProvider(uri)
		logger.Info().Str("url", uri).Msg("using HTTP output provider")
	default:
		return nil, fmt.Errorf("unsuported output provider: %s", scheme)
	}
	return provider, nil
}
