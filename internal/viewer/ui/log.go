package ui

import (
	tui "github.com/marcusolsson/tui-go"
)

type (
	Log struct {
		box  *tui.Box
		logs *tui.Grid
	}
)

func (l *Log) SetTitle(logTitle string) {
	l.box.SetTitle(logTitle)
}

func (l *Log) Draw(logs []string) {
	for _, log := range logs {
		label := tui.NewLabel(log)
		l.logs.AppendRow(label)
	}
}

// newLog Returns a UI component for displaying logs
func newLog() *Log {
	logViewer := tui.NewVBox()
	logViewer.SetTitle("log")
	logViewer.SetBorder(true)
	logViewer.SetSizePolicy(tui.Expanding, tui.Expanding)
	logs := tui.NewGrid(1, 0)
	logViewer.Append(logs)
	return &Log{
		box:  logViewer,
		logs: logs,
	}

}
