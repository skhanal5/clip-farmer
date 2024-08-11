package fetch

import (
	"fmt"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Retrieve content from different media sources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fetch called")
	},
}

func Init() *cobra.Command {
	return fetchCmd
}
