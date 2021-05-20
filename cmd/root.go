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
	"github.com/Casper-Mars/concise-cli/pkg/project"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var moduleName string
var parentVersion string
var CliVersion = "0.1.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "concise-cli",
	Short: "快速构建基础服务项目",
	Long:  `生成基础服务项目，配置需要的文件`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: doTheWork,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&moduleName, "moduleName", "m", "", "指定模块名称，用于pom文件的artifactId")
	rootCmd.PersistentFlags().StringVarP(&parentVersion, "parentVersion", "p", "", "指定父工程的版本")
	rootCmd.Flags().BoolP("version", "v", false, "显示版本号")
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

func doTheWork(cmd *cobra.Command, args []string) {
	showVersion, _ := cmd.Flags().GetBool("version")
	if showVersion {
		fmt.Println("concise-cli:" + CliVersion)
		return
	}

	if moduleName == "" {
		fmt.Println("缺少参数[-m]")
		os.Exit(0)
	}
	if parentVersion == "" {
		fmt.Println("缺少参数[-p]")
		os.Exit(0)
	}
	var builder project.Builder
	builder = project.NewProjectBuilder(project.BuildModeOffline)
	builder.BuildProject("", moduleName, parentVersion)

}
