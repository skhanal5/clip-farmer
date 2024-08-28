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
	} else {
		log.Printf("Successfully deleted all videos in directory: %s", inputDir)
	}
	
}

func deleteAllWithDurationFilter(duration int) {
	dir, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	var count int
	for _, file := range dir {
		filePath := inputDir + "/" + file.Name()
		if (validateFileDuration(filePath, duration)) {
			count += 1
		}
	}
	log.Printf("Successfully deleted %d videos that were <= %d seconds", count, duration)

}

func validateFileDuration(filePath string, duration int) bool {
	fileDuration := getDuration(filePath)
	
	if fileDuration > float64(duration) {
		return false
	}

	time.Sleep(100 * time.Millisecond)
	return removeFile(filePath)
}

func getDuration(filePath string) float64 {
	ctx, cancelFn := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFn()
	data, err := ffprobe.ProbeURL(ctx, filePath)
	if err != nil {
		log.Fatalf("Error probing file: %v", err)
	}
	return data.Format.DurationSeconds
}

func removeFile(filePath string) bool {
	err := os.Remove(filePath)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}