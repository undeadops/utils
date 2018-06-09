package lib

import (
	"log"
	"os"
)

var (
	logger *log.Logger
)

type Logger struct {
	logger   *log.Logger
	logLevel uint
}

func NewLogger() *Logger {
	ll := os.Getenv("LOG_LEVEL")
	l := &Logger{}
	l.setLogLevel(ll)

	l.logger = log.New(os.Stdout, "", log.LstdFlags)
	return l
}

func (l *Logger) setLogLevel(ll string) {
	switch ll {
	case "error":
		l.logLevel = 1
	case "info", "warn":
		l.logLevel = 2
	case "debug":
		l.logLevel = 3
	default:
		l.logLevel = 2
	}
}

func (l *Logger) LogInfo(format string, v ...interface{}) {
	if l.logLevel >= 2 {
		l.setInfo().Printf(format, v...)
	}
}

func (l *Logger) LogError(format string, v ...interface{}) {
	if l.logLevel >= 1 {
		l.setError().Printf(format, v...)
	}
}

func (l *Logger) LogFatal(format string, v ...interface{}) {
	l.setFatal().Fatalf(format, v...)
}

func (l *Logger) LogWarn(format string, v ...interface{}) {
	if l.logLevel >= 2 {
		l.setWarn().Printf(format, v...)
	}
}

func (l *Logger) setFatal() *log.Logger {
	l.logger.SetOutput(os.Stderr)
	l.logger.SetPrefix("FATAL: ")
	return l.logger
}

func (l *Logger) setError() *log.Logger {
	l.logger.SetOutput(os.Stderr)
	l.logger.SetPrefix("ERROR: ")
	return l.logger
}

func (l *Logger) setInfo() *log.Logger {
	l.logger.SetOutput(os.Stdout)
	l.logger.SetPrefix("INFO: ")
	return l.logger
}

func (l *Logger) setWarn() *log.Logger {
	l.logger.SetOutput(os.Stdout)
	l.logger.SetPrefix("WARN: ")
	return l.logger
}
