package logging

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Configure logger level and output format
func Configure(output, level string, pretty bool, sentryDSN string, daemon bool) error {
	if !isatty.IsTerminal(os.Stdout.Fd()) && !daemon && output == "" {
		output = "output.log"
		pretty = false
	}
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
