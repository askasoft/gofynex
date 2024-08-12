package uilog

import (
	"fyne.io/fyne/v2/widget"
	"github.com/askasoft/pango/log"
)

type ConsoleWidget interface {
	WriteText(s string, imp widget.Importance)
}

type ConsoleLogWriter struct {
	log.LogFilter
	log.LogFormatter

	Console ConsoleWidget
}

// Write write message in console.
func (clw *ConsoleLogWriter) Write(le *log.Event) (err error) {
	if clw.Console == nil || clw.Reject(le) {
		return
	}

	bs := clw.Format(le)

	imp := LogLevelToImportance(le.Level)
	clw.Console.WriteText(string(bs), imp)
	return
}

func LogLevelToImportance(lvl log.Level) (imp widget.Importance) {
	switch lvl {
	case log.LevelTrace:
		imp = widget.LowImportance
	case log.LevelDebug:
		imp = widget.MediumImportance
	case log.LevelInfo:
		imp = widget.HighImportance
	case log.LevelWarn:
		imp = widget.WarningImportance
	case log.LevelError:
		imp = widget.DangerImportance
	}
	return
}

// Flush implementing method. empty.
func (clw *ConsoleLogWriter) Flush() {
}

// Close implementing method. empty.
func (clw *ConsoleLogWriter) Close() {
}
