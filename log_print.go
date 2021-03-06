package go_graylog

import (
	"fmt"
	"time"
)

type LogPrint struct{}

var colorLevels = map[string]string{
	"default":  "\033[1;37m%s\033[0m",
	"error":    "\033[1;31m%s\033[0m",
	"critical": "\033[1;31m%s\033[0m",
	"info":     "\033[1;32m%s\033[0m",
	"warning":  "\033[1;33m%s\033[0m",
	"notice":   "\033[1;33m%s\033[0m",
	"debug":    "\033[1;33m%s\033[0m",
}

func (l LogPrint) Debug(mess string, context Context) error {
	return l.log("debug", mess, context)
}

func (l LogPrint) Info(mess string, context Context) error {
	return l.log("info", mess, context)
}

func (l LogPrint) Notice(mess string, context Context) error {
	return l.log("notice", mess, context)
}

func (l LogPrint) Warning(mess string, context Context) error {
	return l.log("warning", mess, context)
}

func (l LogPrint) Error(mess string, context Context) error {
	return l.log("error", mess, context)
}

func (l LogPrint) Critical(mess string, context Context) error {
	return l.log("critical", mess, context)
}

func (l LogPrint) Alert(mess string, context Context) error {
	return l.log("alert", mess, context)
}

func (l LogPrint) Emergency(mess string, context Context) error {
	return l.log("emergency", mess, context)
}

func (l LogPrint) log(level string, mess string, context Context) error {
	date := time.Now().Format("01-02-2006 15:04:05")
	contextStr := ""
	if len(context) > 0 {
		contextStr = fmt.Sprintf(", Context: %v", context)
	}

	logMess := fmt.Sprintf("%s %s: %s%s \n", date, level, mess, contextStr)
	l.formatMess(level, logMess)
	return nil
}

func (l LogPrint) formatMess(level string, mess string) {

	fStr, ok := colorLevels[level]
	if ok {
		fmt.Printf(fStr, mess)
	} else {
		fmt.Println(mess)
	}
}
