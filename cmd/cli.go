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
	"github.com/Casper-Mars/concise-cli/pkg/service"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile       string
	repo          = ""
	branch        = ""
	dist          = ""
	name          = "demo"
	version       = "0.1.0"
	parentVersion = "1.0.17"
	domain        = ""
)

func outputError(error string) {
	fmt.Printf("🚫 %s:\n%s\n", color.RedString("Error"), error)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "concise-cli",
	Short: "项目生成工具",
	Long:  `使用项目模板生成项目，默认使用：`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: createProject,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&repo, "repo", "r", "", "指定模板参考地址")
	rootCmd.Flags().StringVarP(&branch, "branch", "b", "", "指定分支")
	rootCmd.Flags().StringVarP(&dist, "target", "t", "demo-project", "指定生成的目标目录")
	rootCmd.Flags().StringVarP(&name, "name", "n", "demo", "指定生成项目的名称，等于__project_name")
	rootCmd.Flags().StringVarP(&version, "version", "v", "0.1.0", "指定生成项目的版本，等于__project_version")
	rootCmd.Flags().StringVarP(&parentVersion, "parent-version", "p", "1.0.17", "指定生成项目的父工程版本，等于__project_parent_version")
	rootCmd.Flags().StringVarP(&domain, "domain", "d", "", "指定生成项目的域名，等于__project_domain")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".concise-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".concise-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
func createProject(cmd *cobra.Command, args []string) {
	if repo == "" {
		outputError("Repo is missing. Please specify one")
		return
	}
	// make sure the dist is not exist or is empty
	if _, err := os.Stat(dist); !os.IsNotExist(err) {
		outputError(fmt.Sprintf("%s already exist", color.RedString(dist)))
		return
	}
	fmt.Printf("🚀 Creating project with template: %s. Please wait a moment.\n\n", color.GreenString(repo))
	err := service.CreateProject(service.ModeOnline,
		service.WithUrl(repo),
		service.WithBranch(branch),
		service.WithDist(dist),
		service.WithName(name),
		service.WithVersion(version),
		service.WithParentVersion(parentVersion),
		service.WithDomain(domain),
	)
	if err != nil {
		outputError(err.Error())
		return
	}
	fmt.Printf("🍺 Project creation succeeded: %s\n", color.YellowString(dist))
	fmt.Println("🤝 Thanks for using concise cli")
}
