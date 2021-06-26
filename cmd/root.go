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
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	unsplash "github.com/HarshaVardhanNakkina/go-wallpaper/download/unsplash"
)

var cfgFile string
var Resolution string
var Tag string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-wallpaper",
	Short: "Set wallpaper from command line",
	Long: `go-wallpaper is a cli tool for setting wallpapers from different sources through command line
Inspired by https://github.com/thevinter/styli.sh
Currently the image sources is/are Unsplash, Reddit
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return unsplash.DownloadFromUnsplash(Resolution, Tag)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-wallpaper.yaml)")
	rootCmd.Flags().StringVarP(&Resolution, "resolution", "r", "1920x1080", "Resolution of wallpaper")
	rootCmd.Flags().StringVarP(&Tag, "tag", "t", "", "Tag(s)/Search Term(s) comma separated, spaces between (empty by default)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".go-wallpaper" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-wallpaper")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
