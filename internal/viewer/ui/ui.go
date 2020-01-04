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
	u.children.logStream.selectDefaultItem()
	err := u.root.Run()
	if err != nil {
		panic(err)
	}
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

func (u *UI) handleKeyEnter() {
	stream := u.children.logStream.List().SelectedItem()
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

func (u *UI) bind() {
	u.root.SetKeybinding("Esc", func() { u.root.Quit() })
	u.root.SetKeybinding("Down", func() { u.handleKeyDown() })
	u.root.SetKeybinding("Up", func() { u.handleKeyUp() })
	u.root.SetKeybinding("Enter", func() { u.handleKeyEnter() })
}

func (u *UI) SetFetchLog(fetchLog func(ctx context.Context, logStream string) ([]string, error)) {
	u.fetchLog = fetchLog
}

// New Returns a UI reference
func New(streams []string) (*UI, error) {
	wrap := tui.NewHBox()
	logStream := newLogStream(streams)
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
