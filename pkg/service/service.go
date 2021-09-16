package service

import (
	"context"
	"github.com/Casper-Mars/concise-cli/pkg/repo"
	"github.com/Casper-Mars/concise-cli/pkg/subs"
)

type mode int

const (
	ModeOnline mode = iota
	ModeOffline
)

type Option func(o *option)

type option struct {
	mode          mode
	url           string
	dist          string
	branch        string
	name          string
	version       string
	parentVersion string
	domain        string
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

func WithName(name string) Option {
	return func(o *option) {
		o.name = name
	}
}

func WithVersion(version string) Option {
	return func(o *option) {
		o.version = version
	}
}

func WithParentVersion(version string) Option {
	return func(o *option) {
		o.parentVersion = version
	}
}

func WithDomain(domain string) Option {
	return func(o *option) {
		o.domain = domain
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
	case ModeOnline:
		return createOnlineProject(o)
	case ModeOffline:
		panic("implement offline")
	}
	return nil
}

func createOnlineProject(opt *option) error {
	newRepo := repo.NewRepo(opt.url, opt.dist, repo.WithBranch(opt.branch))
	err := newRepo.Clone(context.Background())
	if err != nil {
		return err
	}
	return subs.NewSubsChain(opt.dist, subs.NewDefaultWorkerFactory(),
		subs.WithName(opt.name),
		subs.WithVersion(opt.version),
		subs.WithDomain(opt.domain),
		subs.WithParentVersion(opt.parentVersion),
	).Do(context.Background())
}
