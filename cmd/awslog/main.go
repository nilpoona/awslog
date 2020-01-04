package main

import (
	"context"
	"fmt"

	"github.com/nilpoona/awslog/internal/viewer"
)

type (
	fetcherMock struct{}
)

func (f fetcherMock) Fetch(ctx context.Context, logGroupName string) ([]string, error) {
	return []string{
		`{"level":"Info","message":"foobar"}`,
		`{"level":"Error","message":"foobar"}`,
		`{"level":"Debug","message":"foobar"}`,
		`{"level":"Critical","message":"foobar"}`,
		`{"level":"Warn","message":"foobar"}`,
	}, nil
}

func main() {
	var items []string
	for i := 0; i < 20; i++ {
		items = append(items, fmt.Sprintf("Button %d", i))
	}

	v, err := viewer.New(items, fetcherMock{})
	if err != nil {
		panic(err)
	}

	v.Run(context.Background())
}
