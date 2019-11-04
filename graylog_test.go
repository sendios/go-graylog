package go_graylog

import (
	"testing"
)

func TestSetMaxLevelLogging(t *testing.T) {
	graylog := GrayLog{}
	graylog.SetMaxLevelLogging(LOG_CRIT)

	if graylog.maxLevelLogging != LOG_CRIT {
		t.Error(
			"For", LOG_CRIT,
			"Expected", LOG_CRIT,
			"got", graylog.maxLevelLogging,
		)
	}
}

func TestIsCanWrite(t *testing.T) {
	graylog := GrayLog{}

	graylog.SetMaxLevelLogging(LOG_ERR)
	if graylog.isCanWrite(LOG_INFO) {
		t.Error(
			"Is can write, but min level is",
			LOG_ERR,
		)
	}

	graylog.SetMaxLevelLogging(LOG_ERR)
	if !graylog.isCanWrite(LOG_ALERT) {
		t.Error(
			"Is can't write, but min level is",
			LOG_ERR,
		)
	}
}
