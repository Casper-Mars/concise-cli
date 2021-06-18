/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var projectName string
var parentVersion string
var dependence []string

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "生成业务系统项目",
	Long: `使用时，必须指定父工程的版本、工程项目名称。
还可以指定项目需要用到的外部依赖的组件，例如：mysql、redis。目前支持的外部依赖组件由：mysql、redis、rabbitmq和minio
`,
	Run: createProject,
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().StringVarP(&parentVersion, "parent", "p", "", "指定父项目的版本号")
	projectCmd.Flags().StringVarP(&projectName, "name", "n", "", "指定项目名称")
	projectCmd.Flags().StringArrayVarP(&dependence, "dependence", "d", []string{}, "指定需要使用的外部依赖")
}

func createProject(cmd *cobra.Command, args []string) {
	/*校验参数，工程项目名称和父工程版本号不能为空*/
	if parentVersion == "" {
		fmt.Println("父工程版本号不能为空")
		return
	}
	if projectName == "" {
		fmt.Println("项目工程名称不能为空")
		return
	}

}
