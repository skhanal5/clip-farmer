package post

import (
	"fmt"
	"github.com/spf13/cobra"
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post content onto different media sources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("post called")
	},
}

func Init() *cobra.Command {
	return postCmd
}
