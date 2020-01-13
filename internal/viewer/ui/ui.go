package ui

import (
	"context"
	"sync"

	tui "github.com/marcusolsson/tui-go"
)

type (
	children struct {
		logStream *LogStream
		log       *Log
	}
	// UI a struct that controls the UI component
	UI struct {
		fetchLog   func(ctx context.Context, logStream string) ([]string, error)
		logStreams []string
		root       tui.UI
		children   children
		logs       []string
		mux        sync.Mutex
	}
)

// Run Start execution of UI thread
func (u *UI) Run() {
	err := u.root.Run()
	if err != nil {
		panic(err)
	}
}

func (u *UI) handleKeyDown() {
	u.children.logStream.Select(MoveDown)
}

func (u *UI) handleKeyUp() {
	u.children.logStream.Select(MoveUp)
}

func (u *UI) handleKeyEnter() {
	stream := u.children.logStream.SelectedStreamName()
	u.children.log.SetTitle(stream)
	ctx := context.Background()
	logs, err := u.fetchLog(ctx, stream)
	if err != nil {
		// TODO: Log
	}
	u.mux.Lock()
	u.children.log.Draw(logs)
	u.mux.Unlock()
}

func (u *UI) handleKeyTab() {
	if u.children.logStream.IsFocused() {
		u.children.logStream.OutOfFocus()
		u.children.log.Focused()
		return
	}

	if u.children.log.IsFocused() {
		u.children.log.OutOfFocus()
		u.children.logStream.Focused()
		return
	}
}

func (u *UI) bind() {
	u.root.SetKeybinding("Esc", func() { u.root.Quit() })
	u.root.SetKeybinding("Down", func() { u.handleKeyDown() })
	u.root.SetKeybinding("Up", func() { u.handleKeyUp() })
	u.root.SetKeybinding("Enter", func() { u.handleKeyEnter() })
	u.root.SetKeybinding("Tab", func() { u.handleKeyTab() })
}

func (u *UI) SetFetchLog(fetchLog func(ctx context.Context, logStream string) ([]string, error)) {
	u.fetchLog = fetchLog
}

// New Returns a UI reference
func New(streams []string) (*UI, error) {
	wrap := tui.NewHBox()
	logStream := newLogStream(streams)
	logStream.Focused()
	logViewer := newLog()
	wrap.Append(logStream.Box())
	wrap.Append(logViewer.box)
	logs := []string{}
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
		logs: logs,
		mux:  sync.Mutex{},
	}
	u.bind()
	return u, nil
}
