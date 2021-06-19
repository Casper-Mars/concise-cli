package dir

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
)

type projectDir struct {
	Name  string       `yaml:"name" json:"name"`
	Child []projectDir `yaml:"child" json:"child"`
}

//Build 构建目录
func Build(dirTree []byte, rootPath string) error {
	result := projectDir{}
	err := yaml.Unmarshal(dirTree, &result)
	if err != nil {
		return errors.Wrap(err, "phase dir tree error")
	}
	err = buildDirTress(result, rootPath)
	if err != nil {
		return errors.Wrap(err, "build dir error")
	}
	return nil
}

//buildDirTress 递归构建目录树
func buildDirTress(dir projectDir, parent string) error {
	path := parent + "/" + dir.Name
	err := os.Mkdir(path, 0755|os.ModeDir)
	if err != nil {
		return err
	}
	child := dir.Child
	if len(child) == 0 {
		return nil
	}
	for _, value := range child {
		err := buildDirTress(value, path)
		if err != nil {
			return err
		}
	}
	return nil
}
