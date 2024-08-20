package edit

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a video using ffmpeg",
	Long:  `Edit videos `,
}

func Init() *cobra.Command {
	return editCmd
}
