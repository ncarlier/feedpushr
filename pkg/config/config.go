package config

import (
	"time"
)

var (
	// ListenAddr of the HTTP service
	ListenAddr = FlagEnvString("addr", "HTTP service address", ":8080")
	// PublicURL used for PSHB subscriptions
	PublicURL = FlagEnvString("public-url", "Public URL used for PSHB subscriptions", "")
	// DB location
	DB = FlagEnvString("db", "Database location", "boltdb://data.db")
	// LogLevel (debug/info/warn/error)
	LogLevel = FlagEnvString("log-level", "Logging level", "info")
	// LogPretty writes log using text format
	LogPretty = FlagEnvBool("log-pretty", "Writes log using plain text format", false)
	// Delay is the delay between aggregations
	Delay = FlagEnvDuration("delay", "Delay between aggregations", 1*time.Minute)
	// Timeout is the aggregation timeout
	Timeout = FlagEnvDuration("timeout", "Aggregation timeout", 5*time.Second)
	// CacheRetention is the duration of the cache retention
	CacheRetention = FlagEnvDuration("cache-retention", "Cache retention duration", 72*time.Hour)
	// SentryDSN is the Sentru DSN URL
	SentryDSN = FlagEnvString("sentry-dsn", "Sentry DSN URL", "")
	// ImportFilename is the OPML file to import at boot time
	ImportFilename = FlagString("import", "Import a OPML file at boot time", "")
	// ClearCache is a flag to clear the cache
	ClearCache = FlagBool("clear-cache", "Clear cache at bootstrap", false)
	// Version is a flag to display the version
	Version = FlagBool("version", "Show version", false)
	// Outputs destinations
	Outputs = FlagEnvArray("output", "Output destination", []string{"stdout://"})
	// Plugins to load
	Plugins = FlagEnvArray("plugin", "Plugin to load", []string{})
	// Filters to load
	Filters = FlagEnvArray("filter", "Plugin to load", []string{})
)
