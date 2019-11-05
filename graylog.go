package go_graylog

import (
	"fmt"
	"gopkg.in/Graylog2/go-gelf.v1/gelf"
	"os"
	"time"
)

type GrayLog struct {
	writer   *gelf.Writer
	hostName string
	version  string
	codebase string
}

func (logger *GrayLog) Init(url string, port int, codebase string) error {
	logger.version = "1.0"
	hostName, err := os.Hostname()
	if err != nil {
		return err
	}

	logger.hostName = hostName
	logger.codebase = codebase

	addr := fmt.Sprintf("%s:%d", url, port)
	logger.writer, err = gelf.NewWriter(addr)
	if err != nil {
		return err
	}

	return nil
}

func (logger GrayLog) Debug(mess string, context Context) error {
	return logger.log(LOG_DEBUG, mess, context)
}

func (logger GrayLog) Info(mess string, context Context) error {
	return logger.log(LOG_INFO, mess, context)
}

func (logger GrayLog) Notice(mess string, context Context) error {
	return logger.log(LOG_NOTICE, mess, context)
}

func (logger GrayLog) Warning(mess string, context Context) error {
	return logger.log(LOG_WARNING, mess, context)
}

func (logger GrayLog) Error(mess string, context Context) error {
	return logger.log(LOG_ERR, mess, context)
}

func (logger GrayLog) Critical(mess string, context Context) error {
	return logger.log(LOG_CRIT, mess, context)
}

func (logger GrayLog) Alert(mess string, context Context) error {
	return logger.log(LOG_ALERT, mess, context)
}

func (logger GrayLog) Emergency(mess string, context Context) error {
	return logger.log(LOG_EMERG, mess, context)
}

func (logger GrayLog) Write(p []byte) (n int, err error) {
	mess := string(p)

	return len(mess), logger.Info(mess, Context{})
}

func (logger GrayLog) log(level int32, mess string, context Context) error {

	if _, ok := context["codebase"]; !ok {
		context["codebase"] = logger.codebase
	}

	message := gelf.Message{
		Version:  logger.version,
		Host:     logger.hostName,
		Short:    mess,
		TimeUnix: float64(time.Now().Unix()),
		Level:    level,
		Extra:    context,
	}

	err := logger.writer.WriteMessage(&message)

	if err != nil {
		return err
	}

	return nil
}
