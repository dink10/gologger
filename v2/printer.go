package gologger

import (
	"time"
)

func Print(printerType string, logType string, location string, message interface{}, callInfo *callerInfo, time time.Time, hash string) {
	switch printerType {
	case "console":
		ConsolePrinter(logType, location, message, callInfo, time, hash)
	case "file":
		FilePrinter(logType, location, message, callInfo, time, hash)
	}
}
