package logn

import (
	"os"

	"github.com/theritikchoure/logx"
)

func Info(message string, format ...interface{}) {
	logfLevel("INFO", message, logx.FGGREEN, format...)
}

func Debug(message string, format ...interface{}) {
	logfLevel("DEBUG", message, logx.FGBLUE, format...)
}

func Raw(message string, format ...interface{}) {
	logf(message, "", format...)
}

func Error(message string, format ...interface{}) {
	logfLevel("ERROR", message, logx.FGRED, format...)
}

func Verbose(message string, format ...interface{}) {
	logfLevel("VERBOSE", message, logx.FGYELLOW, format...)
}

// Print error and exit program
func Fatal(message string, format ...interface{}) {
	logfLevel("FATAL", message, logx.FGRED, format...)
	os.Exit(0)
}
