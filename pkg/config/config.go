package config

import (
	"flag"
	"time"
)

// Config contain global configuration
type Config struct {
	ListenAddr     string
	PublicURL      string
	DB             string
	Delay          time.Duration
	Timeout        time.Duration
	CacheRetention time.Duration
	Outputs        ArrayFlags
	Plugins        ArrayFlags
	Filters        ArrayFlags
	ImportFilename string
	ClearCache     bool
	Daemon         bool
	ShowVersion    bool
	LogPretty      bool
	LogLevel       string
	LogOutput      string
	SentryDSN      string
}

var config = Config{}

func init() {
	setFlagEnvString(&config.ListenAddr, "addr", "HTTP service address", ":8080")
	setFlagEnvString(&config.PublicURL, "public-url", "Public URL used for PSHB subscriptions", "")
	setFlagEnvString(&config.DB, "db", "Database location", "boltdb://data.db")
	setFlagEnvString(&config.LogOutput, "log-output", "Log output (STDOUT if empty)", "")
	setFlagEnvString(&config.LogLevel, "log-level", "Logging level (debug, info, warn, error)", "info")
	setFlagEnvBool(&config.LogPretty, "log-pretty", "Writes log using plain text format", false)
	setFlagEnvDuration(&config.Delay, "delay", "Delay between aggregations", 1*time.Minute)
	setFlagEnvDuration(&config.Timeout, "timeout", "Aggregation timeout", 5*time.Second)
	setFlagEnvDuration(&config.CacheRetention, "cache-retention", "Cache retention duration", 72*time.Hour)
	setFlagEnvString(&config.SentryDSN, "sentry-dsn", "Sentry DSN URL", "")
	setFlagString(&config.ImportFilename, "import", "Import a OPML file at boot time", "")
	setFlagBool(&config.Daemon, "daemon", "Start service as daemon", false)
	setFlagBool(&config.ClearCache, "clear-cache", "Clear cache at bootstrap", false)
	setFlagBool(&config.ShowVersion, "version", "Show version", false)
	setFlagEnvArray(&config.Outputs, "output", "Output destination", []string{"stdout://"})
	setFlagEnvArray(&config.Plugins, "plugin", "Plugin to load", []string{})
	setFlagEnvArray(&config.Filters, "filter", "Plugin to load", []string{})

	// set shorthand parameters
	const shorthand = " (shorthand)"
	usage := flag.Lookup("addr").Usage + shorthand
	flag.StringVar(&config.ListenAddr, "l", config.ListenAddr, usage)
	usage = flag.Lookup("daemon").Usage + shorthand
	flag.BoolVar(&config.Daemon, "d", config.Daemon, usage)
	usage = flag.Lookup("version").Usage + shorthand
	flag.BoolVar(&config.ShowVersion, "v", config.ShowVersion, usage)
}

// Get global configuration
func Get() Config {
	return config
}
