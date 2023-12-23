package log

import (
	"testing"
)

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
