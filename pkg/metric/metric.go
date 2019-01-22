package metric

import (
	"expvar"
	"runtime"
	"time"
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

// Configure madditional metrics
func Configure() {
	expvar.Publish("goroutines", expvar.Func(goroutines))
	expvar.Publish("uptime", expvar.Func(uptime))
}
