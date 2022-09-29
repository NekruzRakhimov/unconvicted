package utils

import (
	"crypto/sha1"
	"fmt"
	"runtime"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

func FuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func GenerateHash(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
