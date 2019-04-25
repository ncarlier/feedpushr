package config

import (
	"expvar"
	"time"
)

var conf = expvar.NewMap("config")

func getConfString(val *string) func() interface{} {
	return func() interface{} {
		if val == nil {
			return nil
		}
		return *val
	}
}

func getConfDur(val *time.Duration) func() interface{} {
	return func() interface{} {
		if val == nil {
			return nil
		}
		return val.String()
	}
}

func init() {
	conf.Set("addr", expvar.Func(getConfString(ListenAddr)))
	conf.Set("db", expvar.Func(getConfString(DB)))
	conf.Set("public-url", expvar.Func(getConfString(PublicURL)))
	conf.Set("delay", expvar.Func(getConfDur(Delay)))
	conf.Set("timeout", expvar.Func(getConfDur(Timeout)))
	conf.Set("cache-retention", expvar.Func(getConfDur(CacheRetention)))
}
