package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/askasoft/pango/cog/linkedlist"
	"github.com/askasoft/pango/str"
)

type RichItem struct {
	Text       string
	Importance widget.Importance
}

type Console struct {
	list  *widget.List
	data  linkedlist.LinkedList[RichItem]
	limit int
}

func NewConsole(limit int) *Console {
	c := &Console{
		limit: limit,
	}

	c.list = widget.NewList(c.items, c.createItem, c.updateItem)
	c.list.HideSeparators = true
	return c
}

func (c *Console) Widget() fyne.CanvasObject {
	return c.list
}

func (c *Console) items() int {
	return c.data.Len()
}

func (c *Console) createItem() fyne.CanvasObject {
	lbl := widget.NewLabel("")
	lbl.Wrapping = fyne.TextWrapBreak
	return lbl
}

func (c *Console) updateItem(i widget.ListItemID, o fyne.CanvasObject) {
	item := c.data.Get(i)

	lbl := o.(*widget.Label)
	lbl.Importance = item.Importance
	lbl.SetText(item.Text)

	c.list.SetItemHeight(i, lbl.MinSize().Height)
}

func (c *Console) WriteText(s string, imp widget.Importance) {
	if c.data.Len() > c.limit {
		c.data.PollHead()
	}

	c.data.Push(RichItem{
		Text:       str.StripRight(s),
		Importance: imp,
	})

	c.list.ScrollToBottom()
}

func (c *Console) Clear() {
	c.data.Clear()
	c.list.Refresh()
}
