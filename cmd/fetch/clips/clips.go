package clips

import (
	"fmt"
	"github.com/spf13/cobra"
)

// clipsCmd represents the clips command
var clipsCmd = &cobra.Command{
	Use:   "clips",
	Short: "Fetching clips from different media sources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clips called")
	},
}

func Init() *cobra.Command {
	return clipsCmd
}
