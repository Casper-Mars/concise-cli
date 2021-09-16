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
	cfgFile string
	repo    = ""
	branch  = ""
	dist    = ""
)

func outputError(format string, arg ...string) {
	fmt.Printf("🚫 %s:\n%s\n", color.RedString("Error"), fmt.Sprintf(format, arg))
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
	rootCmd.Flags().StringVarP(&dist, "dist", "d", "demo-project", "指定生成的目标目录")
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
		outputError("%s already exist", color.RedString(dist))
		return
	}
	fmt.Printf("🚀 Creating project with template: %s. Please wait a moment.\n\n", color.GreenString(repo))
	err := service.CreateProject(service.MODE_ONLINE, service.WithUrl(repo), service.WithBranch(branch), service.WithDist(dist))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("🍺 Project creation succeeded: %s\n", color.YellowString(dist))
	fmt.Println("🤝 Thanks for using concise cli")
}
