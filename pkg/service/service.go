package service

import (
	"context"
	"github.com/Casper-Mars/concise-cli/pkg/repo"
)

type mode int

const (
	MODE_ONLINE mode = iota
	MODE_OFFLINE
)

type Option func(o *option)

type option struct {
	mode   mode
	url    string
	dist   string
	branch string
}

func WithUrl(url string) Option {
	return func(o *option) {
		o.url = url
	}
}

func WithDist(dist string) Option {
	return func(o *option) {
		o.dist = dist
	}
}

func WithBranch(branch string) Option {
	return func(o *option) {
		o.branch = branch
	}
}

func CreateProject(mode mode, opts ...Option) error {
	o := &option{
		mode: mode,
	}
	for _, opt := range opts {
		opt(o)
	}
	switch o.mode {
	case MODE_ONLINE:
		return createOnlineProject(o)
	case MODE_OFFLINE:
		panic("implement offline")
	}
	return nil
}

func createOnlineProject(opt *option) error {
	newRepo := repo.NewRepo(opt.url, repo.WithBranch(opt.branch), repo.WithDist(opt.dist))
	return newRepo.Clone(context.Background())
}
