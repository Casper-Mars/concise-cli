package project

import (
	"github.com/Casper-Mars/concise-cli/pkg/creator"
	"os"
)

type OfflineProjectBuilder struct {
}

func (receiver *OfflineProjectBuilder) BuildProject(basePath, moduleName, parentVersion string) {
	if basePath == "" {
		basePath = "."
	}
	dst := basePath + string(os.PathSeparator) + moduleName
	initDir(dst)

	initPomFile(dst, moduleName, parentVersion)
	initMakefile(dst)
	initGitlabFile(dst)
	initGitIgnoreFile(dst)
	initMysqlShell(dst + "/hack")
}

func initDir(basePath string) {
	err := os.Mkdir(basePath, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	path := []string{}
	path = append(path, basePath+"/src/main/java")
	path = append(path, basePath+"/src/main/resources")
	path = append(path, basePath+"/src/test/java")
	path = append(path, basePath+"/src/test/java/unit")
	path = append(path, basePath+"/src/test/java/integration")
	path = append(path, basePath+"/src/test/resources")
	path = append(path, basePath+"/db/test/migration")
	path = append(path, basePath+"/hack")
	err = createDir(path)
	if err != nil {
		err := os.Remove(basePath)
		if err != nil {
			panic(err.Error())
		}
	}
}

func createDir(path []string) error {
	for _, k := range path {
		err := os.MkdirAll(k, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func initGitIgnoreFile(basePath string) {
	err := creator.CreateIgnoreFile(basePath)
	if err != nil {
		panic(err)
	}
}

func initGitlabFile(basePath string) {
	err := creator.CreateGitlabCi(basePath)
	if err != nil {
		panic(err)
	}
}

func initMakefile(basePath string) {
	err := creator.CreateMakefile(basePath)
	if err != nil {
		panic(err)
	}
}

func initPomFile(basePath, moduleName, parentVersion string) {
	err := creator.CreatePom(basePath, moduleName, parentVersion)
	if err != nil {
		panic(err)
	}
}

func initMysqlShell(basePath string) {
	err := creator.CreateMysqlShell(basePath)
	if err != nil {
		panic(err)
	}
}
