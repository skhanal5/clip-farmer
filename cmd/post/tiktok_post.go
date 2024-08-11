/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package post

import (
	"errors"
	"github.com/skhanal5/clip-farmer/manager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	filePath string
)

// tiktokCmd represents the tiktok command
var tiktokCmd = &cobra.Command{
	Use:   "tiktok",
	Short: "Post short-form content onto TikTok",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager, err := buildManager()

		if err != nil {
			return err
		}

		if filePath != "" {
			manager.UploadVideo(filePath)
		}
		return nil
	},
}

func init() {
	postCmd.AddCommand(tiktokCmd)
	tiktokCmd.Flags().StringVarP(&filePath, "path", "p", "",
		"Path to the file containing the media that we want to post onto TikTok.")
}

func buildManager() (manager.TikTokManager, error) {
	clientOAuth := viper.GetString("secrets.tiktok.client-oauth")

	if clientOAuth == "" {
		return manager.TikTokManager{}, errors.New("tiktok client-oauth is not configured")
	}

	return manager.InitTikTokManager(clientOAuth), nil
}
