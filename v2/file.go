package gologger

import (
	"fmt"
	"os"
	"path"
	"time"
)

func FilePrinter(logType string, location string, message interface{}, info *callerInfo, time time.Time, hash string) {
	logFileName := location
	if logFileName == "" {
		logFileName = "log.txt"
	}
	basePath := path.Dir(logFileName)
	filePath := path.Base(logFileName)

	if os.MkdirAll(basePath, 0777) != nil {
		panic("Unable to create directory")
	}

	logFileName = path.Join(basePath, filePath)
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	var logString string
	if info != nil {
		if len(hash) != 0 {
			logString = fmt.Sprintf(
				"[%s] [%s] [%v] [%s::%s::%s] [%d] %s\n",
				logType,
				time.Format("2006-01-02 15:04:05"),
				hash,
				info.packageName,
				info.fileName,
				info.funcName,
				info.line,
				message,
			)
		} else {
			logString = fmt.Sprintf(
				"[%s] [%s] [%s::%s::%s] [%d] %s\n",
				logType,
				time.Format("2006-01-02 15:04:05"),
				info.packageName, info.fileName, info.funcName, info.line, message,
			)
		}
	} else {

		if len(hash) != 0 {
			logString = fmt.Sprintf("[%s] [%s] [%v] %s\n", logType, time.Format("2006-01-02 15:04:05"), hash, message)
		} else {
			logString = fmt.Sprintf("[%s] [%s] %s\n", logType, time.Format("2006-01-02 15:04:05"), message)
		}
	}
	_, fileWriteErr := file.WriteString(logString)
	if fileWriteErr != nil {
		fmt.Println(fileWriteErr)
		os.Exit(1)
	}
	if logType == "FATAL" {
		os.Exit(1)
	}
}
