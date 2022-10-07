package log

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/op/go-logging"
	"os"
)

type Log struct {
	log *logging.Logger
}

var format = logging.MustStringFormatter(
	`%{time:2006/01/02 15:04:05} %{shortfile} [%{level}] %{message}`,
)

func NewLogger(name string, logDir ...string) *Log {
	lg := logging.MustGetLogger(name)
	lg.ExtraCalldepth = 1

	if len(logDir) > 0 {
		fileName := fmt.Sprintf("%s/%s.log", logDir[0], name)
		lb := logging.NewLogBackend(&lumberjack.Logger{
			Filename: fileName,
			MaxSize:  200,  // megabytes
			Compress: true, // disabled by default
		}, "", 0)

		backend1 := logging.NewBackendFormatter(lb, format)
		showConsoleLog := os.Getenv("STDLOG") != "0"
		backends := []logging.Backend{backend1}
		if showConsoleLog {
			backend2 := logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), format)
			backends = append(backends, backend2)
		}
		lg.SetBackend(logging.MultiLogger(backends...))
	}
	return &Log{
		log: lg,
	}
}

func (l *Log) Debug(args ...interface{}) {
	l.log.Debug(args)
}

func (l *Log) Info(args ...interface{}) {
	l.log.Info(args)
}

func (l *Log) Error(args ...interface{}) {
	l.log.Error(args)
}

func (l *Log) Warning(args ...interface{}) {
	l.log.Warning(args)
}
