package go_graylog

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

type Logger struct {
	loggerWrites []LoggerWriter
}

func (l *Logger) AddLoggerWriter(w LoggerWriter) {
	l.loggerWrites = append(l.loggerWrites, w)
}

func (l Logger) Debug(mess string, context Context) {
	for _, w := range l.loggerWrites {
		_ = w.Debug(mess, context)
	}
}

func (l Logger) Info(mess string, context Context) {
	for _, w := range l.loggerWrites {
		_ = w.Info(mess, context)
	}
}

func (l Logger) Notice(mess string, context Context) {
	for _, w := range l.loggerWrites {
		_ = w.Notice(mess, context)
	}
}

func (l Logger) Warning(mess string, context Context) {
	for _, w := range l.loggerWrites {
		_ = w.Warning(mess, context)
	}
}

func (l Logger) Error(mess string, context Context) {
	for _, w := range l.loggerWrites {
		_ = w.Error(mess, context)
	}
}

func (l Logger) Critical(mess string, context Context) {
	for _, w := range l.loggerWrites {
		_ = w.Critical(mess, context)
	}
}

func (l Logger) Alert(mess string, context Context) {
	for _, w := range l.loggerWrites {
		_ = w.Alert(mess, context)
	}
}

func (l Logger) Emergency(mess string, context Context) {
	for _, w := range l.loggerWrites {
		_ = w.Emergency(mess, context)
	}
}
