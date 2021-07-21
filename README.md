# go-graylog

Example:

```
package main

import (
	"fmt"
	go_graylog "github.com/sendios/go-graylog"
)



func main() {
	logger := go_graylog.Logger{}
	
	defer func() {
    	   f r := recover(); r != nil {
    		stack := debug.Stack()
    	    logger.Critical(fmt.Sprintf("Fatal: %s", r), go_graylog.Context{"trace": string(stack)})
    	   }
        }()

	fmtLogger := go_graylog.LogPrint{}
	logger.AddWriter(fmtLogger, go_graylog.LogDebug)

	grayLog := go_graylog.GrayLog{}
	err := grayLog.Init("graylog.sendios", 12201, "test")
	if err != nil {
		fmt.Println(err)
	} else {
		logger.AddWriter(&grayLog, go_graylog.LogErr)
	}


	logger.Info("test_go", go_graylog.Context{
		"something": "some",
		"_file": "file.log",
		"_line": 34,
	})
}
```
