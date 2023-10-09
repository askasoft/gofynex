package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Console struct {
	*container.Scroll

	MaxLines int

	vbox *fyne.Container
}

func NewConsole(maxLines int) *Console {
	c := &Console{
		MaxLines: maxLines,
	}

	c.vbox = container.NewVBox()
	c.Scroll = container.NewVScroll(c.vbox)

	c.ExtendBaseWidget(c)

	return c
}

func (c *Console) Add(s string, imp widget.Importance) {
	if len(c.vbox.Objects) > c.MaxLines {
		c.vbox.Remove(c.vbox.Objects[0])
	}

	lbl := widget.NewLabel(s)
	lbl.Importance = imp
	lbl.Wrapping = fyne.TextWrapBreak
	c.vbox.Add(lbl)
}

func (c *Console) Clear() {
	c.vbox.RemoveAll()
}
