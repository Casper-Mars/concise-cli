package main

import (
	"flag"
	"fmt"
	"github.com/Casper-Mars/concise-cli/paper"
	"os"
)

var moduleName string
var parentVersion string

func main() {
	flag.StringVar(&moduleName, "m", "", "指定模块名称，用于pom文件的artifactId")
	flag.StringVar(&parentVersion, "p", "", "指定父工程的版本")
	flag.Parse()
	if moduleName == "" {
		fmt.Println("缺少参数[-m]")
		flag.Usage()
		os.Exit(0)
	}
	if parentVersion == "" {
		fmt.Println("缺少参数[-p]")
		flag.Usage()
		os.Exit(0)
	}
	initDir(moduleName)
	initFile(moduleName)
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

func initFile(basePath string) {
	initPomFile(basePath, basePath, parentVersion)
	initMakefile(basePath)
	initGitlabFile(basePath)
	initGitIgnoreFile(basePath)
}

func initGitIgnoreFile(basePath string) {
	err := paper.CreateIgnoreFile(basePath)
	if err != nil {
		panic(err)
	}
}

func initGitlabFile(basePath string) {
	err := paper.CreateGitlabCi(basePath)
	if err != nil {
		panic(err)
	}
}

func initMakefile(basePath string) {
	err := paper.CreateMakefile(basePath)
	if err != nil {
		panic(err)
	}
}

func initPomFile(basePath, moduleName, parentVersion string) {
	err := paper.CreatePom(basePath, moduleName, parentVersion)
	if err != nil {
		panic(err)
	}
}
