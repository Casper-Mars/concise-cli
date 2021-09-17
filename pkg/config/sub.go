package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type SubConfig struct {
	Name          []string `yaml:"__project_name"`
	Version       []string `yaml:"__project_version"`
	Domain        []string `yaml:"__project_domain"`
	ParentVersion []string `yaml:"__project_parent_version"`
}

func NewSubConfig(file string) (SubConfig, error) {
	conf := SubConfig{}
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		file, fileReadErr := ioutil.ReadFile(file)
		if fileReadErr != nil {
			return conf, fileReadErr
		}
		unmarshalErr := yaml.Unmarshal(file, &conf)
		if unmarshalErr != nil {
			return conf, unmarshalErr
		}
	}
	return conf, nil
}
