package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var Logger = zerolog.New(
	zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
).With().Timestamp().Caller().Logger()

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SetLoggerLevel(verbose bool) {
	if verbose {
		Logger = Logger.Level(zerolog.TraceLevel)
	} else {
		Logger = Logger.Level(zerolog.NoLevel)
	}
}
