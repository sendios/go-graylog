package go_graylog

//import (
//	"testing"
//)
//
//func TestSetMaxLevelLogging(t *testing.T) {
//	graylog := GrayLog{}
//	graylog.SetMaxLevelLogging(LogCrit)
//
//	if graylog.maxLevelLogging != LogCrit {
//		t.Error(
//			"For", LogCrit,
//			"Expected", LogCrit,
//			"got", graylog.maxLevelLogging,
//		)
//	}
//}
//
//func TestIsCanWrite(t *testing.T) {
//	graylog := GrayLog{}
//
//	graylog.SetMaxLevelLogging(LogErr)
//	if graylog.isCanWrite(LogInfo) {
//		t.Error(
//			"Is can write, but min level is",
//			LogErr,
//		)
//	}
//
//	graylog.SetMaxLevelLogging(LogErr)
//	if !graylog.isCanWrite(LogAlert) {
//		t.Error(
//			"Is can't write, but min level is",
//			LogErr,
//		)
//	}
//}
