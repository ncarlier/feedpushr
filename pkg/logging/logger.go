package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Configure logger level and output format
func Configure(level string, pretty bool) {
	zerolog.TimeFieldFormat = ""
	l := zerolog.InfoLevel
	switch level {
	case "debug":
		l = zerolog.DebugLevel
	case "warn":
		l = zerolog.WarnLevel
	case "error":
		l = zerolog.ErrorLevel
	}
	zerolog.SetGlobalLevel(l)
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	if pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}
