package helpers

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

var (
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func LogError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %v\n", fn, line, err)
}

func Dump(i interface{}) {
	fmt.Printf("%+v\n", i)
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
