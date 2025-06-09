package uilog

import (
	"github.com/askasoft/pango/log"
)

var _ log.Writer = &ConsoleLogWriter{}
