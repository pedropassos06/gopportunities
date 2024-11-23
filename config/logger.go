package config

import (
	"io"
	"log"
	"os"
)

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Err(v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Errf(format string, v ...interface{})
}

type AppLogger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *AppLogger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &AppLogger{
		log.New(writer, "DEBUG: ", logger.Flags()),
		log.New(writer, "INFO: ", logger.Flags()),
		log.New(writer, "WARNING: ", logger.Flags()),
		log.New(writer, "ERROR: ", logger.Flags()),
		writer,
	}
}

// Create non-formatted logs
func (l *AppLogger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *AppLogger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *AppLogger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *AppLogger) Err(v ...interface{}) {
	l.err.Println(v...)
}

// Create formatted logs
func (l *AppLogger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *AppLogger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *AppLogger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *AppLogger) Errf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
