package go_graylog

import (
	"time"
)

const (
	LOG_EMERG   = int32(0)
	LOG_ALERT   = int32(1)
	LOG_CRIT    = int32(2)
	LOG_ERR     = int32(3)
	LOG_WARNING = int32(4)
	LOG_NOTICE  = int32(5)
	LOG_INFO    = int32(6)
	LOG_DEBUG   = int32(7)
)

type Context map[string]interface{}

type LoggerWriter interface {
	Debug(mess string, context Context) error
	Info(mess string, context Context) error
	Notice(mess string, context Context) error
	Warning(mess string, context Context) error
	Error(mess string, context Context) error
	Critical(mess string, context Context) error
	Alert(mess string, context Context) error
	Emergency(mess string, context Context) error
}

type LoggerItem struct {
	writer      LoggerWriter
	maxLogLevel int32
}

func (logger LoggerItem) isCanWrite(level int32) bool {
	return logger.maxLogLevel >= level
}

type Logger struct {
	loggerWrites []LoggerItem
}

func (l *Logger) AddLoggerWriter(w LoggerWriter, maxLogLevel int32) {
	l.loggerWrites = append(l.loggerWrites, LoggerItem{w, maxLogLevel})
}

func (l Logger) Debug(mess string, context Context) {
	l.write(mess, context, LOG_DEBUG)
}

func (l Logger) Info(mess string, context Context) {
	l.write(mess, context, LOG_INFO)
}

func (l Logger) Notice(mess string, context Context) {
	l.write(mess, context, LOG_NOTICE)
}

func (l Logger) Warning(mess string, context Context) {
	l.write(mess, context, LOG_WARNING)
}

func (l Logger) Error(mess string, context Context) {
	l.write(mess, context, LOG_ERR)
}

func (l Logger) Critical(mess string, context Context) {
	l.write(mess, context, LOG_CRIT)
}

func (l Logger) Alert(mess string, context Context) {
	l.write(mess, context, LOG_ALERT)
}

func (l Logger) Emergency(mess string, context Context) {
	l.write(mess, context, LOG_EMERG)
}

func (l Logger) write(mess string, context Context, level int32) {
	m, c := l.prepareData(mess, context)
	for _, w := range l.loggerWrites {
		l.writeToLoggerItem(w, m, c, level)
	}
}

func (l Logger) prepareData(mess string, context Context) (string, Context) {
	m := time.Now().Format("01-02-2006 15:04:05") + mess
	return m, context
}

func (l Logger) writeToLoggerItem(loggerItem LoggerItem, mess string, context Context, level int32) {

	if !loggerItem.isCanWrite(level) {
		return
	}

	switch level {
	case LOG_CRIT:
		_ = loggerItem.writer.Critical(mess, context)
	case LOG_WARNING:
		_ = loggerItem.writer.Warning(mess, context)
	case LOG_EMERG:
		_ = loggerItem.writer.Emergency(mess, context)
	case LOG_ALERT:
		_ = loggerItem.writer.Alert(mess, context)
	case LOG_NOTICE:
		_ = loggerItem.writer.Notice(mess, context)
	case LOG_DEBUG:
		_ = loggerItem.writer.Debug(mess, context)
	case LOG_INFO:
		_ = loggerItem.writer.Info(mess, context)
	}
}
