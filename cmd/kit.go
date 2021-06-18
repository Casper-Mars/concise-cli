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

var kitName string
var conciseParentVersion string

// kitCmd represents the kit command
var kitCmd = &cobra.Command{
	Use:   "kit",
	Short: "生成基础库项目",
	Long:  ``,
	Run:   createKit,
}

func init() {
	rootCmd.AddCommand(kitCmd)
	kitCmd.Flags().StringVarP(&conciseParentVersion, "parent", "p", "", "指定父项目的版本号")
	kitCmd.Flags().StringVarP(&kitName, "name", "n", "", "指定项目名称")
}

func createKit(cmd *cobra.Command, args []string) {
	/*校验参数，工程项目名称和父工程版本号不能为空*/
	if conciseParentVersion == "" {
		fmt.Println("父工程版本号不能为空")
		return
	}
	if kitName == "" {
		fmt.Println("项目工程名称不能为空")
		return
	}
}
