package config

import (
	myerror "github.com/Casper-Mars/concise-cli/pkg/error"
	"github.com/pkg/errors"
)

type projectConfig struct {
	ParentVersion    string
	ParentGroupId    string
	ParentArtifactId string
	Name             string
	Dependence       []string
}

func NewProjectConfig() *projectConfig {
	return &projectConfig{}
}

func (receiver projectConfig) Check() error {
	if receiver.Name == "" {
		return errors.Wrap(myerror.ErrConfigMissing, "项目工程名称不能为空")
	}
	if receiver.ParentVersion == "" {
		return errors.Wrap(myerror.ErrConfigMissing, "父工程版本号不能为空")
	}
	return nil
}
