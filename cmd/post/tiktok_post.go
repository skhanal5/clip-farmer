package post

import (
	"errors"
	"log"

	"github.com/skhanal5/clip-farmer/internal/tiktok"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	filePath string
	dirPath  string
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
		if dirPath != "" {
			manager.UploadVideos(dirPath)
		}
		return nil
	},
}

func init() {
	postCmd.AddCommand(tiktokCmd)
	tiktokCmd.Flags().StringVarP(&filePath, "file", "f", "",
		"Path to the file containing the video that we want to post onto TikTok.")
	tiktokCmd.Flags().StringVarP(&dirPath, "directory", "d", "",
		"Path to the directory with all the videos that we want to post onto TikTok.")
	tiktokCmd.MarkFlagsMutuallyExclusive("file", "directory")
}

func buildManager() (tiktok.TikTokManager, error) {
	log.Println("Building TikTok Manager")
	clientOAuth := viper.GetString("secrets.tiktok.client-oauth")
	if clientOAuth == "" {
		return tiktok.TikTokManager{}, errors.New("tiktok client-oauth is not configured")
	}

	return tiktok.InitTikTokManager(clientOAuth), nil
}
