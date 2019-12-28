package viewer

import (
	"context"

	"github.com/nilpoona/awslog/internal/viewer/ui"
)

type (
	// LogFetcher　interface to get cloud watch log
	LogFetcher interface {
		Fetch(ctx context.Context, logGroupName string) ([]string, error)
	}
	// Viewer struct to control log viewer
	Viewer struct {
		fetcher LogFetcher
		ui      *ui.UI
	}
)

// Run Start execution of UI thread
func (v *Viewer) Run() error {
	return v.ui.Run()
}

// New Returns Viewer
func New(streams []string, fetcher LogFetcher) (*Viewer, error) {
	uiComponent, err := ui.New(streams)
	if err != nil {
		return nil, err
	}
	return &Viewer{
		fetcher: fetcher,
		ui:      uiComponent,
	}, nil
}
