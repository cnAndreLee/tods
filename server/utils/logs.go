package utils

import (
	"fmt"
)

const (
	INFO  = "[Info]"
	WARN  = "[Warn]"
	ERROR = "[Error]"
)

func sendlog(logType string, s string) {
	fmt.Println(logType + " " + s)
}

func LogINFO(s string) {
	sendlog(INFO, s)
}

func LogWARN(s string) {
	sendlog(WARN, s)
}

func LogERROR(s string) {
	sendlog(ERROR, s)
}

// func Sprintf(format string, a ...interface{}) string {
//     p := newPrinter()
//     p.doPrintf(format, a)
//     s := string(p.buf)
//     p.free()
//     return s
// }
