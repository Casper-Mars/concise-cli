package main

import (
	"flag"
	"fmt"
	"github.com/Casper-Mars/concise-cli/pkg/project"
	"os"
)

var moduleName string
var parentVersion string
var cliVersion = "0.1.0"

func main() {
	flag.StringVar(&moduleName, "m", "", "指定模块名称，用于pom文件的artifactId")
	flag.StringVar(&parentVersion, "p", "", "指定父工程的版本")
	showVersion := flag.Bool("v", false, "显示版本号")
	online := flag.Bool("o", false, "是否使用在线的模板")
	flag.Parse()
	if showVersion != nil && *showVersion {
		fmt.Println("concise-cli:" + cliVersion)
		return
	}
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
	var builder project.Builder
	if online != nil && *online {
		builder = project.NewProjectBuilder(project.BuildModeOnline)
	} else {
		builder = project.NewProjectBuilder(project.BuildModeOffline)
	}
	builder.BuildProject("", moduleName, parentVersion)
}
