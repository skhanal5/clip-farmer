package clean

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/vansante/go-ffprobe.v2"
)

var (
	inputDir string
	duration int
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up files generated from either fetching or editing",
	Run: func(cmd *cobra.Command, args []string) {
		if duration != 0 {
			deleteAllWithDurationFilter(duration)
		} else {
			deleteAll(inputDir)
		}
	},
}

func Init() *cobra.Command {
	cleanCmd.Flags().StringVarP(&inputDir, "directory", "d", "", "Directory containing the mp4 files that we would like to delete")
	cleanCmd.Flags().IntVarP(&duration, "duration", "s", 0, "The duration in seconds to filter mp4 files. All files less than this will be deleted")
	cleanCmd.MarkFlagRequired("directory")
	return cleanCmd
}

func deleteAll(inputDir string) {
	err := os.RemoveAll(inputDir)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteAllWithDurationFilter(duration int) {
	dir, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	for _, file := range dir {
		file, err := os.Open(file.Name())

		if err != nil {
			log.Panicf("Error opening file: %v", err)
		}

		data, err := ffprobe.ProbeReader(ctx, file)
		if err != nil {
			log.Panicf("Error probing file: %v", err)
		}

		if (data.Format.DurationSeconds <= float64(duration)) {
			os.Remove(file.Name())
		}
	}

}