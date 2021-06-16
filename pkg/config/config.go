package config

import (
	"time"
)

// Config contain global configuration
type Config struct {
	ListenAddr         string        `flag:"listen-addr" desc:"HTTP listen address" default:":8080"`
	PublicURL          string        `flag:"public-url" desc:"Public URL used for PSHB subscriptions" default:""`
	DB                 string        `flag:"db" desc:"Database location" default:"boltdb://data.db"`
	Delay              time.Duration `flag:"delay" desc:"Delay between aggregations" default:"1m"`
	Timeout            time.Duration `flag:"timeout" desc:"Aggregation timeout" default:"5s"`
	CacheRetention     time.Duration `flag:"cache-retention" desc:"Cache retention duration" default:"72h"`
	FanOutDelay        time.Duration `flag:"fan-out-delay" desc:"Delay between deployment of each aggregator" default:"0s"`
	Plugins            []string      `flag:"plugin" desc:"Plugin to load" default:""`
	ImportFilename     string        `flag:"import" desc:"Import an OPML file at service startup" default:""`
	ClearCache         bool          `flag:"clear-cache" desc:"Clear cache at service startup" default:"false"`
	ClearConfig        bool          `flag:"clear-config" desc:"Clear configuration at service startup" default:"false"`
	LogPretty          bool          `flag:"log-pretty" desc:"Writes log using plain text format" default:"false"`
	LogLevel           string        `flag:"log-level" desc:"Logging level (debug, info, warn or error)" default:"info"`
	LogOutput          string        `flag:"log-output" desc:"Log output (STDOUT if empty)" default:""`
	Authn              string        `flag:"authn" desc:"Authentication method (Basic HTTP with password file, OIDC issuer URL or none)" default:".htpasswd"`
	AuthorizedUsername string        `flag:"authorized-username" desc:"Authorized username" default:"*"`
	SentryDSN          string        `flag:"sentry-dsn" desc:"Sentry DSN URL" default:""`
	ExploreProvider    string        `flag:"explore-provider" desc:"Provider used to find RSS feeds" default:"default"`
	ServiceName        string        `flag:"service-name" desc:"Service name used by the service registry" default:"feedpushr"`
}
