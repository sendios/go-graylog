package go_graylog

import (
	"fmt"
	"runtime/debug"
)

const (
	LogEmerg   = int32(0)
	LogAlert   = int32(1)
	LogCrit    = int32(2)
	LogErr     = int32(3)
	LogWarning = int32(4)
	LogNotice  = int32(5)
	LogInfo    = int32(6)
	LogDebug   = int32(7)
)

type Context map[string]interface{}

type LogWriter interface {
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
	writer      LogWriter
	maxLogLevel int32
}

func (logger LoggerItem) isCanWrite(level int32) bool {
	return logger.maxLogLevel >= level
}

type Logger struct {
	loggerWrites []LoggerItem
}

func (l *Logger) AddWriter(w LogWriter, maxLogLevel int32) {
	l.loggerWrites = append(l.loggerWrites, LoggerItem{w, maxLogLevel})
}

func (l Logger) Debug(mess string, context Context) {
	l.write(mess, context, LogDebug)
}

func (l Logger) Info(mess string, context Context) {
	l.write(mess, context, LogInfo)
}

func (l Logger) Notice(mess string, context Context) {
	l.write(mess, context, LogNotice)
}

func (l Logger) Warning(mess string, context Context) {
	l.write(mess, context, LogWarning)
}

func (l Logger) Error(mess string, context Context) {
	l.write(mess, context, LogErr)
}

func (l Logger) Critical(mess string, context Context) {
	l.write(mess, context, LogCrit)
}

func (l Logger) Alert(mess string, context Context) {
	l.write(mess, context, LogAlert)
}

func (l Logger) Emergency(mess string, context Context) {
	l.write(mess, context, LogEmerg)
}

func (l Logger) write(mess string, context Context, level int32) {
	m, c := l.prepareData(mess, context)
	for _, w := range l.loggerWrites {
		l.writeToLoggerItem(w, m, c, level)
	}
}

func (l Logger) prepareData(mess string, context Context) (string, Context) {
	return mess, context
}

func (l Logger) writeToLoggerItem(loggerItem LoggerItem, mess string, context Context, level int32) {

	if !loggerItem.isCanWrite(level) {
		return
	}

	switch level {
	case LogCrit:
		_ = loggerItem.writer.Critical(mess, context)
	case LogWarning:
		_ = loggerItem.writer.Warning(mess, context)
	case LogErr:
		_ = loggerItem.writer.Error(mess, context)
	case LogEmerg:
		_ = loggerItem.writer.Emergency(mess, context)
	case LogAlert:
		_ = loggerItem.writer.Alert(mess, context)
	case LogNotice:
		_ = loggerItem.writer.Notice(mess, context)
	case LogDebug:
		_ = loggerItem.writer.Debug(mess, context)
	case LogInfo:
		_ = loggerItem.writer.Info(mess, context)
	}
}

func (l Logger) Recover() {
	if r := recover(); r != nil {
		stack := debug.Stack()
		l.Error(fmt.Sprintf("Fatal: %s", r), Context{"trace": string(stack)})
	}
}
