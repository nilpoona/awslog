package ui

import tui "github.com/marcusolsson/tui-go"

type (
	children struct {
		logStream *tui.Box
		log       *tui.Box
	}
	// UI a struct that controls the UI component
	UI struct {
		root     tui.UI
		children children
	}
)

// Run Start execution of UI thread
func (u *UI) Run() error {
	return u.root.Run()
}

// New Returns a UI reference
func New(streams []string) (*UI, error) {
	wrap := tui.NewHBox()
	logStream := newLogStream(streams)
	logViewer := newLog()
	wrap.Append(logStream)
	wrap.Append(logViewer)
	u, err := tui.New(wrap)
	if err != nil {
		return nil, err
	}
	u.SetKeybinding("Esc", func() { u.Quit() })
	return &UI{
		root: u,
		children: children{
			logStream: logStream,
			log:       logViewer,
		},
	}, nil
}
