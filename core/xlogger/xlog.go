package xlogger

import (
	"fmt"
)

const (
	LEVEL_VERBOSE = 0
	LEVEL_INFO    = 1
	LEVEL_DEBUG         = 2
	LEVEL_WARNING       = 3
	LEVEL_ERROR         = 4
	LEVEL_EMPTY         = 10
)

const (
	COLOR_RED    = "\033[31m"
	COLOR_YELLOW = "\033[33m"
	COLOR_WHITE  = "\033[37m"
	COLOR_BLUE   = "\033[34m"
	COLOR_CYAN   = "\033[36m"
	COLOR_GREEN  = "\033[32m"
	RESET        = "\033[0m"
)

type Level int8

var LogLevel Level = LEVEL_VERBOSE
var DefaultTag = ""


func showLog(msg string, levelStatus Level) {
	if LogLevel > levelStatus {
		return
	}

	switch levelStatus {
	case LEVEL_VERBOSE:
		fmt.Println(string(COLOR_GREEN), "🤗  <------------------------------     VERBOSE :    ["+DefaultTag+"] # "+msg, string(RESET))
		return
	case LEVEL_INFO:
		fmt.Println(string(COLOR_CYAN), "😏  <------------------------------     INFO :       ["+DefaultTag+"] # "+msg, string(RESET))
		return
	case LEVEL_DEBUG:
		fmt.Println(string(COLOR_BLUE), "🤕  <------------------------------     DEBUG :      ["+DefaultTag+"] # "+msg, string(RESET))
		return
	case LEVEL_WARNING:
		fmt.Println(string(COLOR_YELLOW), "😵  <------------------------------     WARNING :    ["+DefaultTag+"] # "+msg, string(RESET))
		return
	case LEVEL_ERROR:
		fmt.Println(string(COLOR_RED), "💩  <------------------------------     ERROR :      [" + DefaultTag+"] # "+msg, string(RESET))
		return
	}
}

func Verbose(msg string) {
	showLog(msg, LEVEL_VERBOSE)
}

func Info(msg string) {
	showLog(msg, LEVEL_INFO)
}

func Debug(msg string) {
	showLog(msg, LEVEL_DEBUG)
}
func Err(msg string) {
	showLog(msg, LEVEL_ERROR)
}

func Warn(msg string) {
	showLog(msg, LEVEL_WARNING)
}
