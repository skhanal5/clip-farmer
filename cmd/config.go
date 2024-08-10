package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set the clip-farmer config for environment variables",
	Long: `Override the default loading behavior of the clip-farmer config file by either passing in your own
file to load environment variables from or by setting the environment variables manually through commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileFlag, _ := cmd.Flags().GetString("path")
		fmt.Println(fileFlag)
		if fileFlag != "" {
			loadConfig(fileFlag)
		}
	},
}

func init() {
	configCmd.Flags().StringVarP(&cfgFile, "path", "p", "config.yaml",
		`The location of the yaml file to load configuration from.`)
	rootCmd.AddCommand(configCmd)
}

func loadConfig(cfgFile string) {
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Loading environment variable from config file:", viper.ConfigFileUsed())
	}
}

func SaveConfig() {
	err := viper.WriteConfigAs("config.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to save config to config.yaml")
	}
}
