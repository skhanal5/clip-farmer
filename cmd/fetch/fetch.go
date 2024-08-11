package fetch

import (
	"fmt"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var FetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Retrieve data from different media sources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fetch called")
	},
}

func Init() *cobra.Command {
	return FetchCmd
}
