package fetch

import (
	"github.com/skhanal5/clip-farmer/cmd/fetch/clips"
	"github.com/skhanal5/clip-farmer/cmd/fetch/oauth"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var FetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Retrieve data from different media sources",
}

func Init() *cobra.Command {
	FetchCmd.AddCommand(clips.Init())
	FetchCmd.AddCommand(oauth.Init())
	return FetchCmd
}
