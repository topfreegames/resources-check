// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config *viper.Viper
var json bool

// ConfigFile is the configuration file used for running a command
var ConfigFile string

// Verbose determines how verbose maestro will run under
var Verbose int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "resources-check",
	Short: "Resources Check app",
	Long:  `Ensure every Deployment, Statefulset and Daemonset on Kubernetes has request and limits set.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().BoolVarP(
		&json, "json", "j",
		false, "json output mode")

	RootCmd.PersistentFlags().IntVarP(
		&Verbose, "verbose", "v", 0,
		"Verbosity level => v0: Error, v1=Warning, v2=Info, v3=Debug",
	)

	RootCmd.PersistentFlags().StringVarP(
		&ConfigFile, "config", "c", "./config/local.yaml",
		"config file",
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config = viper.New()
	if ConfigFile != "" { // enable ability to specify config file via flag
		config.SetConfigFile(ConfigFile)
	}
	config.SetConfigType("yaml")
	config.SetEnvPrefix("rcheck")
	config.AddConfigPath(".")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "__"))
	config.AutomaticEnv()

	// If a config file is found, read it in.
	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("Config file %s failed to load: %s.\n", ConfigFile, err.Error())
		panic("Failed to load config file")
	}
}
