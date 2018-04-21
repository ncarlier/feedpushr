package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Configuration specification
type Configuration struct {
	ListenAddr     string        `default:":8080" split_words:"true"`
	PublicURL      string        `split_words:"true"`
	Store          string        `default:"boltdb://data.db"`
	Output         string        `default:"stdout"`
	LogLevel       string        `default:"info" split_words:"true"`
	LogPretty      bool          `default:"false" split_words:"true"`
	Delay          time.Duration `default:"1m"`
	Timeout        time.Duration `default:"5s"`
	CacheRetention time.Duration `default:"72h" split_words:"true"`
	Filters        []string
}

// Config holder
var Config Configuration

func init() {
	err := envconfig.Process("app", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
