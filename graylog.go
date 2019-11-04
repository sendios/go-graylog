package go_graylog

import (
	"fmt"
	"gopkg.in/Graylog2/go-gelf.v1/gelf"
	"os"
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

type GrayLog struct {
	writer          *gelf.Writer
	hostName        string
	version         string
	codebase        string
	maxLevelLogging int32
}

type Context map[string]interface{}

func (logger *GrayLog) Init(url string, port int, codebase string) error {
	logger.version = "1.0"
	hostName, err := os.Hostname()
	if err != nil {
		return err
	}

	logger.hostName = hostName
	logger.codebase = codebase

	if logger.maxLevelLogging == 0 {
		logger.maxLevelLogging = 3
	}

	addr := fmt.Sprintf("%s:%d", url, port)
	logger.writer, err = gelf.NewWriter(addr)
	if err != nil {
		return err
	}

	return nil
}

func (logger *GrayLog) SetMaxLevelLogging(maxLevelLogging int32) {
	logger.maxLevelLogging = maxLevelLogging
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

func (logger GrayLog) isCanWrite(level int32) bool {
	return logger.maxLevelLogging >= level
}

func (logger GrayLog) log(level int32, mess string, context Context) error {

	if !logger.isCanWrite(level) {
		return nil
	}

	if _, ok := context["codebase"]; !ok {
		context["codebase"] = logger.codebase
	}

	message := gelf.Message{
		Version:  logger.version,
		Host:     logger.hostName,
		Short:    mess,
		TimeUnix: float64(time.Now().Unix()),
		Level:    level,
		Facility: "0",
		Extra:    context,
	}

	err := logger.writer.WriteMessage(&message)

	if err != nil {
		return err
	}

	return nil
}
