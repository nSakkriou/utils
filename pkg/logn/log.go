package logn

import (
	"fmt"
	"strings"
	"time"

	"github.com/theritikchoure/logx"
)

func logf(message, color string, format ...interface{}) {
	message = strings.ToLower(message)

	if len(format) == 0 {
		logx.Logf(message, color, "", format...)
	} else {
		formattedMessage := fmt.Sprintf(message, format...)
		logx.Logf(formattedMessage, color, "")
	}
}

func logfLevel(level, message, color string, format ...interface{}) {
	prefixedMessage := fmt.Sprintf("[%s] [%s]\t%s", timestamp(), level, message)
	logf(prefixedMessage, color, format...)
}

func timestamp() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}
