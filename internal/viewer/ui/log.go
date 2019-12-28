package ui

import tui "github.com/marcusolsson/tui-go"

// newLog Returns a UI component for displaying logs
func newLog() *tui.Box {
	logViewer := tui.NewVBox(tui.NewLabel("log ga deruyo"))
	logViewer.SetTitle("log")
	logViewer.SetBorder(true)
	return logViewer
}
