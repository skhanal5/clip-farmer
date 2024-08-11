package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set the clip-farmer config for environment variables",
	Long: `Override the default loading behavior of the clip-farmer config file by either passing in your own
file to load environment variables from or by setting the environment variables manually through commands.`,
}

func Init() *cobra.Command {
	return configCmd
}

func SaveConfig(path string) {
	viper.SetConfigFile(path)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("Failed to write config:", err)
	}
}
