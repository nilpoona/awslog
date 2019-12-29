package ui

import (
	tui "github.com/marcusolsson/tui-go"
)

type (
	// LogStream UI component to display log stream
	LogStream struct {
		root *tui.Box
		list *tui.List
	}
)

// Box
func (l *LogStream) Box() *tui.Box {
	return l.root
}

// List
func (l *LogStream) List() *tui.List {
	return l.list
}

const defaultSelectIndex = 0

// newLogStream Returns a UI component that displays a list of log streams
func newLogStream(streams []string) *LogStream {
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

	return &LogStream{
		root: box,
		list: list,
	}
}
