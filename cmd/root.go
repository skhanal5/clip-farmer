package cmd

import (
	"github.com/skhanal5/clip-farmer/cmd/config"
	"github.com/skhanal5/clip-farmer/cmd/fetch"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clip-farmer",
	Short: "A CLI tool to automatically fetch, edit, and post short form content.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(config.Init())
	rootCmd.AddCommand(fetch.Init())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile("./config.yaml")
	viper.AutomaticEnv() // read in environment variables that match
	_ = viper.ReadInConfig()
}
