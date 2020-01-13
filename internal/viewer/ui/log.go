package ui

import (
	"strings"

	tui "github.com/marcusolsson/tui-go"
	"github.com/tidwall/gjson"
)

type (
	Log struct {
		box        *tui.Box
		scrollArea *tui.ScrollArea
		logs       *tui.Grid
	}
)

func (l *Log) SetTitle(logTitle string) {
	l.box.SetTitle(logTitle)
}

func (l *Log) logLabel(log string) *tui.Label {
	if !gjson.Valid(log) {
		return tui.NewLabel(log)
	}

	result := gjson.Get(log, "level")
	level := result.String()
	if level == "" {
		return tui.NewLabel(log)
	}
	level = strings.ToLower(level)
	label := tui.NewLabel(log)

	switch {
	case strings.Contains(level, "error"):
		label.SetStyleName(themeErrorLabel)
	case strings.Contains(level, "warn"):
		label.SetStyleName(themeWarnLabel)
	case strings.Contains(level, "info"):
		label.SetStyleName(themeInfoLabel)
	case strings.Contains(level, "debug"):
		label.SetStyleName(themeDebugLabel)
	default:
		return label
	}
	label.SetStyleName(themeWarnLabel)
	return label
}

func (l *Log) Draw(logs []string) {
	for _, log := range logs {
		l.logs.AppendRow(l.logLabel(log))
	}
}

// OutOfFocus out of focus
func (l *Log) OutOfFocus() {
	l.scrollArea.SetFocused(false)
}

// Focused Focus on
func (l *Log) Focused() {
	l.scrollArea.SetFocused(true)
}

// IsFocused Determine if focus is on
func (l *Log) IsFocused() bool {
	return l.scrollArea.IsFocused()
}

// Empty Empty the log display area
func (l *Log) Empty() {
	l.logs.RemoveRows()
}

// newLog Returns a UI component for displaying logs
func newLog() *Log {
	logViewer := tui.NewVBox()
	logViewer.SetTitle("log")
	logViewer.SetBorder(true)
	logViewer.SetSizePolicy(tui.Expanding, tui.Expanding)
	logs := tui.NewGrid(1, 0)
	scrollArea := tui.NewScrollArea(logs)
	logViewer.Append(scrollArea)
	return &Log{
		box:        logViewer,
		scrollArea: scrollArea,
		logs:       logs,
	}

}
