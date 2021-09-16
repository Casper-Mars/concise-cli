package subs

import (
	"context"
	"fmt"
	"github.com/Casper-Mars/concise-cli/pkg/config"
	"golang.org/x/sync/errgroup"
)

type Option func(o *options)

type options struct {
	configFilename string
	name           string
	version        string
	domain         string
	parentVersion  string
}

func WithConfigFilename(filename string) Option {
	return func(o *options) {
		o.configFilename = filename
	}
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func WithVersion(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

func WithDomain(domain string) Option {
	return func(o *options) {
		o.domain = domain
	}
}

func WithParentVersion(version string) Option {
	return func(o *options) {
		o.parentVersion = version
	}
}

type Chain struct {
	projectRootPath string
	keyword         map[string]map[string]string
	workerFactory   WorkerFactory
}

func (c Chain) Do(ctx context.Context) error {
	if 0 == len(c.keyword) {
		return nil
	}
	group, _ := errgroup.WithContext(ctx)
	for file, placeholder := range c.keyword {
		worker := c.workerFactory.CreateWorker(fmt.Sprintf("%s/%s", c.projectRootPath, file), placeholder)
		group.Go(func() error {
			return worker.Substitute(context.Background())
		})
	}
	return group.Wait()
}

func (c Chain) init(o *options) {
	// get config
	configFile := fmt.Sprintf("%s/%s", c.projectRootPath, o.configFilename)
	conf, err := config.NewSubConfig(configFile)
	if err == nil {
		setFunc := func(value, valueName string, files []string) {
			for _, k := range files {
				m := c.keyword[k]
				if m == nil {
					c.keyword[k] = make(map[string]string, 0)
					m = c.keyword[k]
				}
				m[valueName] = value
			}
		}
		if o.name != "" {
			setFunc(o.name, "__project_name", conf.Name)
		}
		if o.version != "" {
			setFunc(o.version, "__project_version", conf.Version)
		}
		if o.domain != "" {
			setFunc(o.domain, "__project_domain", conf.Domain)
		}
		if o.parentVersion != "" {
			setFunc(o.parentVersion, "__project_parent_version", conf.ParentVersion)
		}
	}
}

func NewSubsChain(projectRootPath string, workerFactory WorkerFactory, opts ...Option) *Chain {
	option := &options{
		configFilename: "concise.yaml",
	}
	for _, opt := range opts {
		opt(option)
	}
	c := &Chain{
		projectRootPath: projectRootPath,
		keyword:         make(map[string]map[string]string, 0),
		workerFactory:   workerFactory,
	}
	c.init(option)

	return c
}
