package logn

import (
	"os"
)

func Raw(message string, format ...interface{}) {
	if ShowRaw {
		logf(message, "", format...)
	}
}

func Debug(message string, format ...interface{}) {
	if ShowDebug {
		logfLevel(PrefixDebug, message, ColorDebug, format...)
	}
}

func Verbose(message string, format ...interface{}) {
	if ShowVerbose {
		logfLevel(PrefixVerbose, message, ColorVerbose, format...)
	}
}

func Info(message string, format ...interface{}) {
	if ShowInfo {
		logfLevel(PrefixInfo, message, ColorInfo, format...)
	}
}

func Warn(message string, format ...interface{}) {
	if ShowInfo {
		logfLevel(PrefixWarn, message, ColorWarn, format...)
	}
}

func Error(message string, format ...interface{}) {
	if ShowError {
		logfLevel(PrefixError, message, ColorError, format...)
	}
}

// Print error and exit program
func Fatal(message string, format ...interface{}) {
	if ShowFatal {
		logfLevel(PrefixFatal, message, ColorFatal, format...)
	}

	os.Exit(0)
}
