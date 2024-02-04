package config

/*
			Reset: "\033[0m"
			Styles:
				Bold:      "\033[1m"
				Strike:    "\033[9m"
				Italic:    "\033[3m"
				Underline: "\033[4m"
			Colors:
				Red:    "\033[31m"
	//			Blue:   "\033[34m"
				Cyan:   "\033[36m"
				Green:  "\033[32m"
				White:  "\033[37m"
				Yellow: "\033[33m"
				Purple: "\033[30m"
*/

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	err     *log.Logger
	info    *log.Logger
	debug   *log.Logger
	writer  *io.Writer
	warning *log.Logger
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)
	return &Logger{
		err:     log.New(writer, "\033[31m\033[1mERROR\033[0m ", logger.Flags()),
		info:    log.New(writer, "\033[36m\033[1mINFO\033[0m ", logger.Flags()),
		debug:   log.New(writer, "\033[32m\033[1mDEBUG\033[0m ", logger.Flags()),
		warning: log.New(writer, "\033[33m\033[1mWARNING\033[0m ", logger.Flags()),
		writer:  &writer,
	}
}

// Create a Non-Formatted Logs.
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

// Create a Formatted Logs.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
