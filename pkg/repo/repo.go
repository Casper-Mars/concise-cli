package repo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var ErrCloneFail = errors.New("git clone fail")

type Option func(repo *Repo)

type Repo struct {
	url    string
	dist   string
	branch string
}

//WithBranch set branch
func WithBranch(branch string) Option {
	return func(repo *Repo) {
		repo.branch = branch
	}
}

//WithDist set target dir
func WithDist(dist string) Option {
	return func(repo *Repo) {
		repo.dist = dist
	}
}

//NewRepo
func NewRepo(url string, opts ...Option) *Repo {
	r := &Repo{
		url:  strings.TrimSuffix(url, "/"),
		dist: "demo-project",
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (receiver Repo) Clone(ctx context.Context) error {
	var command *exec.Cmd
	if receiver.branch != "" {
		command = exec.Command("git", "clone", "-q", "-b", receiver.branch, receiver.url, receiver.dist)
	} else {
		command = exec.Command("git", "clone", "-q", receiver.url, receiver.dist)
	}
	errOutput := bytes.Buffer{}
	command.Stderr = &errOutput
	err := command.Run()
	if err != nil {
		return fmt.Errorf("%w:%s\n", ErrCloneFail, errOutput.String())
	}
	return nil
}
