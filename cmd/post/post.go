package post

import (
	"github.com/spf13/cobra"
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post content onto different media sources",
}

func Init() *cobra.Command {
	return postCmd
}
