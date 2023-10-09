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
func (cjw *ConsoleLogWriter) Write(le *log.Event) (err error) {
	if cjw.Console == nil || cjw.Reject(le) {
		return
	}

	bs := cjw.Format(le)

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

	cjw.Console.Add(string(bs), imp)
	cjw.Console.ScrollToBottom()
	return
}

// Flush implementing method. empty.
func (cjw *ConsoleLogWriter) Flush() {
}

// Close implementing method. empty.
func (cjw *ConsoleLogWriter) Close() {
}

// Clear clear the output
func (cjw *ConsoleLogWriter) Clear() {
	if cjw.Console != nil {
		cjw.Console.Clear()
	}
}
