package config

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set the clip-farmer config environment variables",
	Long:  `Override the default loading behavior of the clip-farmer config file by setting the environment variables manually.`,
}

func Init() *cobra.Command {
	return configCmd
}

func SaveConfig(path string) {
	viper.SetConfigFile(path)
	err := viper.WriteConfig()
	if err != nil {
		log.Fatal("Failed to write config:", err)
	} else {
		log.Printf("Updated config to path: %s", path)
	}
}
