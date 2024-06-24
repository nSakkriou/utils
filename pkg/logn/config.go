package logn

import "github.com/theritikchoure/logx"

// SHOW
// You can select which level you want print
var (
	ShowRaw     = true
	ShowDebug   = true
	ShowVerbose = true
	ShowInfo    = true
	ShowWarn    = true
	ShowError   = true
	ShowFatal   = true
)

// COLOR
// You can edit level color
// https://talyian.github.io/ansicolors/
//  \x1b[38;2;r;g;bm - foreground
//	\x1b[48;2;r;g;bm - background
var (
	ColorDebug   = logx.FGBLUE
	ColorVerbose = logx.FGYELLOW
	ColorInfo    = logx.FGGREEN
	ColorWarn    = "\x1b[38;5;214m"
	ColorError   = logx.FGRED
	ColorFatal   = logx.FGRED
)

// LEVEL PREFIX
var (
	PrefixDebug   = "debug"
	PrefixVerbose = "verbose"
	PrefixInfo    = "info"
	PrefixWarn    = "warn"
	PrefixError   = "error"
	PrefixFatal   = "fatal"
)

// LOG FORMAT
var (
	Lowercase  = true
	TimeFormat = "2006-01-02 15:04:05"
)
