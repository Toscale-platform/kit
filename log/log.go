package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

var l = log.Output(zerolog.ConsoleWriter{
	Out:        os.Stdout,
	TimeFormat: "02 Jan 15:04:05",
}).With().Caller().Logger()

func init() {
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

/*
l.Panic()
l.Fatal()
l.Error()
l.Warn()
l.Info()
l.Debug()
l.Trace()
*/

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
