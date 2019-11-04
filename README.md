# go-graylog


```
package main

import (
	"fmt"
	"github.com/mailfire/go-graylog"
	"log"
)

func main() {
	logger := go_graylog.GrayLog{}
	err := logger.Init("localhost", 15501, "test")
	if err != nil {
		fmt.Println(err)
		return
	}

	logger.SetMaxLevelLogging(go_graylog.LOG_INFO)
	err = logger.Info("test_go", go_graylog.Context{})

	if err != nil {
		fmt.Println(err)
	}
	//You may use this as
	log.SetOutput(&logger)
}
```
