package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/askasoft/pango/cog"
)

type Console struct {
	*container.Scroll

	MaxLines int

	Segments cog.ArrayList[widget.RichTextSegment]

	RichText *widget.RichText
}

func NewConsole(maxLines int) *Console {
	c := &Console{
		MaxLines: maxLines,
	}

	c.RichText = widget.NewRichText()
	c.RichText.Wrapping = fyne.TextWrapBreak

	c.Scroll = container.NewVScroll(c.RichText)

	return c
}

func (c *Console) Add(s string, imp widget.Importance) {
	if c.Segments.Len() > c.MaxLines {
		c.Segments.PollHead()
	}

	var color fyne.ThemeColorName
	switch imp {
	case widget.LowImportance:
		color = theme.ColorNameDisabled
	case widget.MediumImportance:
		color = theme.ColorNameForeground
	case widget.HighImportance:
		color = theme.ColorNamePrimary
	case widget.DangerImportance:
		color = theme.ColorNameError
	case widget.WarningImportance:
		color = theme.ColorNameWarning
	case widget.SuccessImportance:
		color = theme.ColorNameSuccess
	default:
		color = theme.ColorNameForeground
	}

	c.Segments.Add(&widget.TextSegment{
		Style: widget.RichTextStyle{
			Alignment: fyne.TextAlignLeading,
			ColorName: color,
			Inline:    true,
			TextStyle: fyne.TextStyle{},
		},
		Text: s,
	})
	c.RichText.Segments = c.Segments.Values()
	c.RichText.Refresh()
}

func (c *Console) Refresh() {
	c.RichText.Segments = c.Segments.Values()
	c.RichText.Refresh()
	c.Scroll.Refresh()
}

func (c *Console) Clear() {
	c.Segments.Clear()
	c.Refresh()
}
