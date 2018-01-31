package gologger

import (
	"fmt"
	"time"
)

func ConsolePrinter(logType string, location string, message interface{}, info *callerInfo, time time.Time, hash string) {
	color := getColor(logType, location)
	color.Set()
	if info != nil {
		if len(hash) != 0 {
			fmt.Printf(
				"[%s] [%s] [%v] [%s::%s::%s] [%d] %s\n",
				logType, time.Format("2006-01-02 15:04:05"), hash,
				info.packageName, info.fileName, info.funcName, info.line, message,
			)
		} else {
			fmt.Printf(
				"[%s] [%s] [%s::%s::%s] [%d] %s\n",
				logType, time.Format("2006-01-02 15:04:05"),
				info.packageName, info.fileName, info.funcName, info.line, message,
			)
		}

	} else {
		if len(hash) != 0 {
			fmt.Printf("[%s] [%s] [%v] %s\n", logType, time.Format("2006-01-02 15:04:05"), hash, message)
		} else {
			fmt.Printf("[%s] [%s] %s\n", logType, time.Format("2006-01-02 15:04:05"), message)
		}
	}
	Unset()
}

func getColor(logType, location string) *Color {
	var color *Color

	if location == "simple" {
		color = New(Reset)
		return color
	}

	switch logType {
	case "LOG":
		color = New(Reset)
		break
	case "MSG":
		color = New(FgBlue)
		break
	case "INF":
		color = New(FgGreen)
		break
	case "WRN":
		color = New(FgMagenta)
		break
	case "DBG":
		color = New(FgYellow)
		break
	case "ERR":
		color = New(FgRed)
		break
	case "RSS":
		color = New(Reset)
		break
	default:
		color = New(Reset)
		break
	}
	return color
}
