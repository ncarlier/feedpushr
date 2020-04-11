package server

import (
	"path/filepath"

	"github.com/ncarlier/feedpushr/v3/pkg/config"
	"github.com/ncarlier/feedpushr/v3/pkg/plugin"
	"github.com/rs/zerolog/log"
)

func loadPlugins(conf config.Config) error {
	// Auto add plugins if not configured
	plugins := conf.Plugins
	if len(plugins) == 0 {
		var err error
		plugins, err = filepath.Glob("feedpushr-*.so")
		if err != nil {
			log.Error().Err(err).Msg("unable to autoload plugins")
		}
	}
	// Load plugins
	err := plugin.Configure(plugins)
	if err != nil {
		log.Error().Err(err).Msg("unable to init plugins")
		return err
	}
	return nil
}
