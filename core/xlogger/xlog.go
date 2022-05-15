package xlogger

import (
	"fmt"
)

const (
	VERBOSE = 0
	INFO    = 1
	DEBUG   = 2
	WARNING = 3
	ERROR   = 4
	EMPTY   = 10
)

// log.warn()
// level = error

type Level int8

var LogLevel Level = VERBOSE
var DefaultTag = ""

func Debug(msg string) {
	showLog(msg, DEBUG)
}

func showLog(msg string, levelStatus Level) {
	if LogLevel > levelStatus {
		return
	}

	switch levelStatus {
	case VERBOSE:
		fmt.Println("------------------------------ VERBOSE : " + DefaultTag + " # " + msg)
		return
	case INFO:
		fmt.Println("------------------------------ INFO : " + DefaultTag + " # " + msg)
		return
	case DEBUG:
		fmt.Println("------------------------------ DEBUG : " + DefaultTag + " # " + msg)
		return
	case WARNING:
		fmt.Println("------------------------------ WARNING : " + DefaultTag + " # " +  msg)
		return
	case ERROR:
		fmt.Println("------------------------------ ERROR : " + DefaultTag + " # " + msg)
		return
	}
}

func Verbose(msg string) {
	showLog(msg, VERBOSE)
}

func Info(msg string) {
	showLog(msg, INFO)
}

func Err(msg string) {
	showLog(msg, ERROR)
}

func Warn(msg string) {
	showLog(msg, WARNING)
}
