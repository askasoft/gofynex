package uilog

import (
	"fyne.io/fyne/v2/widget"
	"github.com/askasoft/fynes/widgets"
	"github.com/askasoft/pango/log"
)

type ConsoleLogWriter struct {
	log.LogFilter
	log.LogFormatter

	Console *widgets.Console
}

// Write write message in console.
func (clw *ConsoleLogWriter) Write(le *log.Event) (err error) {
	if clw.Console == nil || clw.Reject(le) {
		return
	}

	bs := clw.Format(le)

	var imp widget.Importance
	switch le.Level {
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

	clw.Console.Add(string(bs), imp)
	clw.Console.ScrollToBottom()
	return
}

// Flush implementing method. empty.
func (clw *ConsoleLogWriter) Flush() {
}

// Close implementing method. empty.
func (clw *ConsoleLogWriter) Close() {
}

// Clear clear the output
func (clw *ConsoleLogWriter) Clear() {
	if clw.Console != nil {
		clw.Console.Clear()
	}
}
