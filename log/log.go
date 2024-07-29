package log

import (
	"github.com/Toscale-platform/kit/env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

var l = log.With().Caller().Logger()

func init() {
	if env.GetBool("DEBUG") {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		l = log.With().Caller().Logger()
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}
}

func Logger() *zerolog.Logger {
	return &l
}

func Panic() *zerolog.Event {
	return l.Panic()
}

func Fatal() *zerolog.Event {
	return l.Fatal()
}

func Error() *zerolog.Event {
	return l.Error()
}

func Warn() *zerolog.Event {
	return l.Warn()
}

func Info() *zerolog.Event {
	return l.Info()
}

func Debug() *zerolog.Event {
	return l.Debug()
}

func Trace() *zerolog.Event {
	return l.Trace()
}
