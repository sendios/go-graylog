package go_graylog

import (
	"fmt"
)

type LogPrint struct{}

func (l LogPrint) Debug(mess string, context Context) {
	l.log("debug", mess, context)
}

func (l LogPrint) Info(mess string, context Context) {
	l.log("info", mess, context)
}

func (l LogPrint) Notice(mess string, context Context) {
	l.log("notice", mess, context)
}

func (l LogPrint) Warning(mess string, context Context) {
	l.log("warning", mess, context)
}

func (l LogPrint) Error(mess string, context Context) {
	l.log("error", mess, context)
}

func (l LogPrint) Critical(mess string, context Context) {
	l.log("critical", mess, context)
}

func (l LogPrint) Alert(mess string, context Context) {
	l.log("alert", mess, context)
}

func (l LogPrint) Emergency(mess string, context Context) {
	l.log("emergency", mess, context)
}

func (l LogPrint) log(level string, mess string, context Context) {
	fmt.Printf("Level: %s,\n Message %s,\n Context: %v \n\n", level, mess, context)
}
