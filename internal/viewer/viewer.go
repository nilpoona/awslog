package viewer

import (
	"context"
	"errors"
	"fmt"

	"github.com/nilpoona/awslog/internal/viewer/ui"
)

type (
	// LogFetcher　interface to get cloud watch log
	LogFetcher interface {
		Fetch(ctx context.Context, logGroupName string) ([]string, error)
	}
	// Viewer struct to control log viewer
	Viewer struct {
		fetch   chan string
		receive chan string
		fetcher LogFetcher
		ui      *ui.UI
	}
	mockLogger  struct{}
	fetchResult struct {
		logs []string
		err  error
	}
)

func (m mockLogger) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args)
	fmt.Printf("\n")
}

func (v *Viewer) FetchLog(ctx context.Context, logStream string) ([]string, error) {
	resultCh := make(chan fetchResult)
	go func() {
		logs, err := v.fetcher.Fetch(ctx, logStream)
		if err != nil {
			resultCh <- fetchResult{
				logs: nil,
				err:  err,
			}
			close(resultCh)
			return
		}
		resultCh <- fetchResult{
			logs: logs,
			err:  nil,
		}
		close(resultCh)
	}()

	for result := range resultCh {
		return result.logs, result.err
	}

	return nil, errors.New("failed to fetch_log")
}

// Run Start execution of UI thread
func (v *Viewer) Run(ctx context.Context) {
	v.ui.Run()
}

// New Returns Viewer
func New(streams []string, fetcher LogFetcher) (*Viewer, error) {
	uiComponent, err := ui.New(streams)
	if err != nil {
		return nil, err
	}
	// tui.SetLogger(mockLogger{})
	v := &Viewer{
		fetcher: fetcher,
	}

	uiComponent.SetFetchLog(v.FetchLog)
	v.ui = uiComponent
	return v, nil
}
