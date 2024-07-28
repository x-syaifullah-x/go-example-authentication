package logger

import (
	_log "log"
)

func Printf(format string, v ...any) {
	_log.Printf(format, v...)
}

func Print(v any) {
	_log.Print(v)
}

func Fatal(message any) {
	_log.Fatal(message)
}
