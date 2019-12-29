package ui

import (
	tui "github.com/marcusolsson/tui-go"
)

type (
	children struct {
		logStream *LogStream
		log       *tui.Box
	}
	// UI a struct that controls the UI component
	UI struct {
		logStreams []string
		root       tui.UI
		children   children
	}
)

// Run Start execution of UI thread
func (u *UI) Run() error {
	return u.root.Run()
}

func (u *UI) handleKeyDown() {
	nextIndex := u.children.logStream.List().Selected() + 1
	if nextIndex > len(u.logStreams)-1 {
		nextIndex = 0
	}
	u.children.logStream.List().Select(nextIndex)
}

func (u *UI) handleKeyUp() {
	nextIndex := u.children.logStream.List().Selected() - 1
	if nextIndex < 0 {
		nextIndex = len(u.logStreams) - 1
	}
	u.children.logStream.List().Select(nextIndex)
}

func (u *UI) bind() {
	u.root.SetKeybinding("Esc", func() { u.root.Quit() })
	u.root.SetKeybinding("Down", func() { u.handleKeyDown() })
	u.root.SetKeybinding("Up", func() { u.handleKeyUp() })
}

// New Returns a UI reference
func New(streams []string) (*UI, error) {
	wrap := tui.NewHBox()
	logStream := newLogStream(streams)
	logViewer := newLog()
	wrap.Append(logStream.Box())
	wrap.Append(logViewer)
	r, err := tui.New(wrap)
	if err != nil {
		return nil, err
	}
	u := &UI{
		logStreams: streams,
		root:       r,
		children: children{
			logStream: logStream,
			log:       logViewer,
		},
	}
	u.bind()
	return u, nil
}
