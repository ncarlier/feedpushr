package metric

import (
	"expvar"
	"runtime"
	"time"

	"github.com/ncarlier/feedpushr/v2/pkg/config"
)

var startTime = time.Now().UTC()

func goroutines() interface{} {
	return runtime.NumGoroutine()
}

// uptime is an expvar.Func compliant wrapper for uptime info.
func uptime() interface{} {
	uptime := time.Since(startTime)
	return int64(uptime)
}

// Configure additional metrics
func Configure() {
	// Export system metrics
	expvar.Publish("goroutines", expvar.Func(goroutines))
	expvar.Publish("uptime", expvar.Func(uptime))
	// Export configuration variables
	config.ExportConfigVars()
}
