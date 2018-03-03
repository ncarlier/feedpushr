package logging

import (
	"context"
	"fmt"

	"github.com/goadesign/goa"
	"github.com/rs/zerolog"
)

// adapter is the zerolog goa adapter logger.
type adapter struct {
	zerolog.Logger
}

// NewLogAdapter wraps a zerolog logger into a goa logger adapter.
func NewLogAdapter(logger zerolog.Logger) goa.LogAdapter {
	return &adapter{Logger: logger}
}

// Logger returns the zerolog logger stored in the given context if any, nil otherwise.
func Logger(ctx context.Context) *zerolog.Logger {
	logger := goa.ContextLogger(ctx)
	if a, ok := logger.(*adapter); ok {
		return &a.Logger
	}
	return nil
}

// Info logs messages using zerolog.
func (a *adapter) Info(msg string, data ...interface{}) {
	a.Logger.Info().Fields(data2fields(data)).Msg(msg)
}

// Error logs errors using zerolog.
func (a *adapter) Error(msg string, data ...interface{}) {
	a.Logger.Error().Fields(data2fields(data)).Msg(msg)
}

// New creates a new logger given a context.
func (a *adapter) New(data ...interface{}) goa.LogAdapter {
	return &adapter{Logger: a.Logger.With().Fields(data2fields(data)).Logger()}
}

func data2fields(keyvals []interface{}) map[string]interface{} {
	n := (len(keyvals) + 1) / 2
	res := make(map[string]interface{}, n)
	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		var v interface{} = goa.ErrMissingLogValue
		if i+1 < len(keyvals) {
			v = keyvals[i+1]
		}
		res[fmt.Sprintf("%v", k)] = v
	}
	return res
}
