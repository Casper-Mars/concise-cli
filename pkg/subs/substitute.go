package subs

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
)

var (
	ErrFileNotExist = errors.New("file not exist")
)

type Worker interface {
	sub(projectRootPath string) error
}

type Chain struct {
	worker          []Worker
	projectRootPath string
}

func (c Chain) Do(ctx context.Context) error {
	if 0 == len(c.worker) {
		return nil
	}
	group, _ := errgroup.WithContext(ctx)
	for _, worker := range c.worker {
		curWorker := worker
		group.Go(func() error {
			return curWorker.sub(c.projectRootPath)
		})
	}
	return group.Wait()
}

func NewSubsChain(projectRootPath string, workers ...Worker) *Chain {
	return &Chain{
		projectRootPath: projectRootPath,
		worker:          workers,
	}
}
