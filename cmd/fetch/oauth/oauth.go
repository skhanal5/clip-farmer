package oauth

import (
	"fmt"
	"github.com/spf13/cobra"
)

// oauthCmd represents the oauth command
var oauthCmd = &cobra.Command{
	Use:   "oauth",
	Short: "Fetching oauth tokens from services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("oauth called")
	},
}

func Init() *cobra.Command {
	return oauthCmd
}
