package config

import (
	"expvar"
)

// ExportConfigVars export some configuration variables to expvar
func ExportConfigVars() {
	conf := Get()
	expvar.NewString("config.addr").Set(conf.ListenAddr)
	expvar.NewString("config.db").Set(conf.DB)
	expvar.NewString("config.public-url").Set(conf.PublicURL)
	expvar.NewString("config.delay").Set(conf.Delay.String())
	expvar.NewString("config.timeout").Set(conf.Timeout.String())
	expvar.NewString("config.cache-retention").Set(conf.CacheRetention.String())
}
