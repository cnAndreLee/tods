package utils

import (
	"fmt"

	"github.com/cnAndreLee/tods_server/config"
)

const (
	INFO  = "[Info]"
	WARN  = "[Warn]"
	ERROR = "[Error]"
	DEBUG = "[Debug]"
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

func LogDebug(s string) {
	if config.CONFIG.IsDebug == true {
		sendlog(DEBUG, s)
	}
}
