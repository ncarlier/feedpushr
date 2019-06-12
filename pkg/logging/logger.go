package logging

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Configure logger level and output format
func Configure(output, level string, pretty bool, sentryDSN string) error {
	out := os.Stdout
	if output != "" {
		var err error
		out, err = os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
	}

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
	var w io.Writer = out
	if pretty {
		w = zerolog.ConsoleWriter{Out: out}
	}
	if sentryDSN != "" {
		w = zerolog.MultiLevelWriter(w, SentryWriter(sentryDSN))
	}
	log.Logger = zerolog.New(w).With().Timestamp().Logger()

	return nil
}
