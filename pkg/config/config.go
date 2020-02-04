package config

import (
	"time"
)

// Config contain global configuration
type Config struct {
	ListenAddr     string        `flag:"listen-addr" desc:"HTTP listen address" default:":8080"`
	PublicURL      string        `flag:"public-url" desc:"Public URL used for PSHB subscriptions" default:""`
	DB             string        `flag:"db" desc:"Database location" default:"boltdb://data.db"`
	Delay          time.Duration `flag:"delay" desc:"Delay between aggregations" default:"1m"`
	Timeout        time.Duration `flag:"timeout" desc:"Aggregation timeout" default:"5s"`
	CacheRetention time.Duration `flag:"cache-retention" desc:"Cache retention duration" default:"72h"`
	Plugins        []string      `flag:"plugin" desc:"Plugin to load" default:""`
	ImportFilename string        `flag:"import" desc:"Import an OPML file at boot time" default:""`
	ClearCache     bool          `flag:"clear-cache" desc:"Clear cache at bootstrap" default:"false"`
	ClearConfig    bool          `flag:"clear-config" desc:"Clear configuration at bootstrap" default:"false"`
	LogPretty      bool          `flag:"log-pretty" desc:"Writes log using plain text format" default:"false"`
	LogLevel       string        `flag:"log-level" desc:"Logging level (debug, info, warn or error)" default:"info"`
	LogOutput      string        `flag:"log-output" desc:"Log output (STDOUT if empty)" default:""`
	SentryDSN      string        `flag:"sentry-dsn" desc:"Sentry DSN URL" default:""`
}
