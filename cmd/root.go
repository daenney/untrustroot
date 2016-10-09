// Copyright © 2016 Daniele Sluijters
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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var keychain string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "untrustroot",
	Short: "Analyze and update a Keychain's trusted certificates",
	Long: `Provides the ability to analyze and trust/untrust certificates
in a given macOS Keychain. It defaults to the keychain containing the System
Root Certificates.

The "analyze" subcommand will take a look at the keychain and print some
information about each CA, including if they have a CRL or support OCSP.`,
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

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "path to a untrustroot config file (searches /etc/untrustroot, XDG_CONFIG_HOME/untrustroot and $HOME/.config/untrustroot)")
	RootCmd.PersistentFlags().StringVarP(&keychain, "keychain", "k", "/System/Library/Keychains/SystemRootCertificates.keychain", "path to the keychain")
	viper.BindPFlag("keychain", RootCmd.PersistentFlags().Lookup("keychain"))
}

func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("untrustroot")
	viper.AddConfigPath("/etc/untrustroot")
	viper.AddConfigPath("$XDG_CONFIG_HOME/untrustroot")
	viper.AddConfigPath("$HOME/.config/untrustroot")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
