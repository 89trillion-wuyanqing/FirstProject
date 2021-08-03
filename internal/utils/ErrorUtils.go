package utils

import (
	"fmt"
	"runtime/debug"
)

func CaptureError(err error) {
	if err != nil {
		//log.Fatal("ERROR:"+err.Error())
		panic(err)
	}
}

func Dispose() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic recover! p:", p)
			str, _ := p.(string)
			fmt.Println(str)
			debug.PrintStack()
		}
	}()
}
