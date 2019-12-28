package ui

import (
	tui "github.com/marcusolsson/tui-go"
)

const defaultSelectIndex = 0

// newLogStream Returns a UI component that displays a list of log streams
func newLogStream(streams []string) *tui.Box {
	box := tui.NewVBox()
	list := tui.NewList()
	var items []string
	for _, stream := range streams {
		items = append(items, stream)
	}
	list.AddItems(items...)
	list.Select(defaultSelectIndex)
	box.SetBorder(true)
	box.SetTitle("log streams")
	streamList := tui.NewScrollArea(list)
	box.Append(streamList)
	box.SetFocused(true)
	return box
}
