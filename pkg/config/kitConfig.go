package config

import (
	myerror "github.com/Casper-Mars/concise-cli/pkg/error"
	"github.com/pkg/errors"
)

type kitConfig struct {
	KitName       string
	ParentVersion string
	Dependence    []string
}

func NewKitConfig() *kitConfig {
	return &kitConfig{}
}

func (receiver kitConfig) Check() error {
	if receiver.ParentVersion == "" {
		return errors.WithMessage(myerror.ErrConfigMissing, "父工程版本号不能为空")
	}
	if receiver.KitName == "" {
		return errors.WithMessage(myerror.ErrConfigMissing, "项目工程名称不能为空")
	}
	return nil
}
