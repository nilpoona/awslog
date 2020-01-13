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
	MoveDirection string
)

var (
	MoveUp   MoveDirection = "Up"
	MoveDown MoveDirection = "Down"
)

func (d MoveDirection) Next(index int) int {
	switch d {
	case MoveUp:
		return index - 1
	case MoveDown:
		return index + 1
	default:
		return index
	}
}

// Box
func (l *LogStream) Box() *tui.Box {
	return l.root
}

// Select
func (l *LogStream) Select(direction MoveDirection) {
	if l.IsFocused() {
		nextIndex := direction.Next(l.list.Selected())
		if nextIndex > l.list.Length()-1 {
			nextIndex = 0
		}
		if nextIndex < 0 {
			nextIndex = l.list.Length() - 1
		}
		l.list.Select(nextIndex)
	}
}

func (l *LogStream) SelectedStreamName() string {
	return l.list.SelectedItem()
}

// OutOfFocused out of focus
func (l *LogStream) OutOfFocus() {
	l.list.SetFocused(false)
}

// Focused Focus
func (l *LogStream) Focused() {
	l.list.SetFocused(true)
}

// IsFocused Determine if focus is on
func (l *LogStream) IsFocused() bool {
	return l.list.IsFocused()
}

// selectDefaultItem
func (l *LogStream) selectDefaultItem() {
	l.list.Select(defaultSelectIndex)
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
	return &LogStream{
		root: box,
		list: list,
	}
}
