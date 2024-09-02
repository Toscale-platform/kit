package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"path/filepath"
	"strconv"
	"testing"
)

func BenchmarkBase(b *testing.B) {
	var l = log.Output(io.Discard).With().Caller().Logger()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	for i := 0; i < b.N; i++ {
		l.Info().Str("foo", "bar").Msg("Hello world")
	}
}

func BenchmarkPretty(b *testing.B) {
	var l = log.Output(zerolog.ConsoleWriter{
		Out:        io.Discard,
		TimeFormat: "02 Jan 15:04",
	}).With().Caller().Logger()

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	for i := 0; i < b.N; i++ {
		l.Info().Str("foo", "bar").Msg("Hello world")
	}
}

func TestLogger(t *testing.T) {
	Logger().Info().Msg("New info message")
}

func TestPanic(t *testing.T) {
	defer func() {
		recover()
	}()

	Panic().Msg("New panic message")
}

func TestFatal(t *testing.T) {
	Fatal().Msg("New fatal message")
}

func TestError(t *testing.T) {
	Error().Msg("New error message")
}

func TestWarn(t *testing.T) {
	Warn().Msg("New warn message")
}

func TestInfo(t *testing.T) {
	Info().Msg("New info message")
}

func TestDebug(t *testing.T) {
	Debug().Msg("New debug message")
}

func TestTrace(t *testing.T) {
	Trace().Msg("New trace message")
}
