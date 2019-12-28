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
	panic("implement me")
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

	if err := v.Run(); err != nil {
		panic(err)
	}
}
