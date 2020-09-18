package helpers

import (
	"fmt"
	"log"
	"runtime"
)

func LogError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %v\n", fn, line, err)
}

func Dump(i interface{}) {
	fmt.Printf("%+v\n", i)
}
