package gologger

import (
	"crypto/sha1"
	"fmt"
)

const (
	CONSOLE    = "console"
	FILE       = "file"
	SimpleLog  = "simple"
	ColoredLog = "color"
)

type GoLogger struct {
	printerType string
	location    string
	enable      bool
	debug       bool
	full        bool
}

func GetLogger(enable bool, debugLog bool, fullLog bool, selector ...string) GoLogger {
	if len(selector) == 0 {
		return GoLogger{printerType: CONSOLE, location: ColoredLog, enable: enable, debug: debugLog, full: fullLog}
	}
	return GoLogger{printerType: selector[0], location: selector[1], enable: enable, debug: debugLog, full: fullLog}
}

func (l *GoLogger) Log(message interface{}, sHash interface{}) {
	if l.debug {
		logPrinter(l.enable, l.full, l.printerType, "LOG", l.location, message, getSum(sHash))
	}
}

func (l *GoLogger) Message(message interface{}, sHash interface{}) {
	if l.debug {
		logPrinter(l.enable, l.full, l.printerType, "MSG", l.location, message, getSum(sHash))
	}
}

func (l *GoLogger) Info(message interface{}, sHash interface{}) {
	if l.debug {
		logPrinter(l.enable, l.full, l.printerType, "INF", l.location, message, getSum(sHash))
	}
}

func (l *GoLogger) Warn(message interface{}, sHash interface{}) {
	if l.debug {
		logPrinter(l.enable, l.full, l.printerType, "WRN", l.location, message, getSum(sHash))
	}
}

func (l *GoLogger) Debug(message interface{}, sHash interface{}) {
	if l.debug {
		logPrinter(l.enable, l.full, l.printerType, "DBG", l.location, message, getSum(sHash))
	}
}

func (l *GoLogger) Error(message interface{}, sHash interface{}) {
	logPrinter(l.enable, l.full, l.printerType, "ERR", l.location, message, getSum(sHash))
}

func (l *GoLogger) Fatal(message interface{}, sHash interface{}) {
	logPrinter(l.enable, l.full, l.printerType, "FATAL", l.location, message, getSum(sHash))
}

func (l *GoLogger) ReplaceMessage(message interface{}, sHash interface{}) {
	if l.debug {
		logPrinter(l.enable, l.full, l.printerType, "RSS", l.location, message, getSum(sHash))
	}
}

func getSum(sHash interface{}) (sum string) {
	h := sha1.New()
	switch hash := sHash.(type) {
	case string:
		if len(hash) != 0 {
			h.Write([]byte(hash))
			sum = fmt.Sprintf("%x", h.Sum(nil))
		}
	}

	return
}