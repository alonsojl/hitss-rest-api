package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func Write(err interface{}) {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		msg := fmt.Sprintf("%s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		fileName := fmt.Sprintf("%s/errors.log", os.Getenv("ROOT"))
		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		pfx := "ERROR: "
		msgLogger := log.New(file, pfx, log.Ldate|log.Ltime|log.Lmsgprefix)
		msgLogger.Println(msg)
	}
}
