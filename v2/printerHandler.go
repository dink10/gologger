package gologger

import (
	"path"
	"runtime"
	"strings"
	"time"
)

type callerInfo struct {
	packageName string
	fileName    string
	funcName    string
	line        int
}

func logPrinter(enable bool, full bool, printerType string, logType string, location string, message interface{}, hash string) {
	if enable {
		info := retrieveCallInfo(full)
		Print(printerType, logType, location, message, info, time.Now(), hash)
	}
}

func retrieveCallInfo(full bool) *callerInfo {
	if full {
		pc, file, line, _ := runtime.Caller(3)
		_, fileName := path.Split(file)
		parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
		pl := len(parts)
		packageName := ""
		funcName := parts[pl-1]

		if parts[pl-2][0] == '(' {
			funcName = parts[pl-2] + "." + funcName
			packageName = strings.Join(parts[0:pl-2], ".")
		} else {
			packageName = strings.Join(parts[0:pl-1], ".")
		}
		return &callerInfo{
			packageName: packageName,
			fileName:    fileName,
			funcName:    funcName,
			line:        line,
		}
	}

	return nil
}
