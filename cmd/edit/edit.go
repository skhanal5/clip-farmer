package edit

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

var (
	inputDir string
	inputPath string
	outputPath string
	blurredOption bool
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a video",
	RunE: func(cmd *cobra.Command, args []string) error {
		if blurredOption {
			if inputPath != "" {
				createVideoWithBlurredBackground(inputPath, outputPath)
			}
			if inputDir != "" {
				createVideosWithBlurredBackground(inputDir, outputPath)
			}
		}		
		return nil
	},
}

func Init() *cobra.Command {
	editCmd.Flags().StringVarP(&inputDir, "directory", "d", "",
	"Directory containing the mp4 files that we would like to edit")

	editCmd.Flags().StringVarP(&inputPath, "file", "f", "",
		"Path of the mp4 file that we would like to edit")
	editCmd.Flags().StringVarP(&outputPath, "output", "o", "",
		"Path of the resulting edited video.")
	

	// One of input or directory is needed to start editing
	// Output is always required
	// Input and directory must not be provided together
	editCmd.MarkFlagsOneRequired("file", "directory")
	editCmd.MarkFlagsMutuallyExclusive("file", "directory")
	editCmd.MarkFlagsOneRequired("output")

	editCmd.Flags().BoolVarP(&blurredOption, "blurred", "b", false,
	"Make a video that is overlayed ontop of a blurred background")
	editCmd.MarkFlagsOneRequired("blurred")
	return editCmd
}

func createVideosWithBlurredBackground(inputDir string, outputPath string) {
	err := os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for _, file := range dir {
		fileInfo, err := file.Info()
		if err != nil {
			log.Fatal(err)
		}

		// do this asynchronously
		wg.Add(1)
		go func () {
			defer wg.Done()
			createVideoWithBlurredBackground(inputDir + "/" + fileInfo.Name(), outputPath)	
		}()
	}

	wg.Wait()
}

func getFilename(inputPath string) string{
	return filepath.Base(inputPath)
}

func createVideoWithBlurredBackground(inputPath string, outputDir string) {
	log.Print("Blurring the video")
	

	err := os.MkdirAll("bin", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Change resolution to TikTok's viewport
	// Apply box blur
	// Remove audio
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-s", "1080x1920", "-vf", "boxblur=50", "-an", "bin/tmp.mp4")
	fmt.Println(cmd.Args)
	_, err = cmd.CombinedOutput()
	
	// delete video using defer
	defer deleteTmpFiles()
	
	if err != nil {
		log.Fatalf("Failed to blur video")
	} 
	
	// then overlay the video with the "clear" version
	log.Print("Overlaying the original video on top of the blurred video.")
	outputFilePath := outputPath + getFilename(inputPath)
	cmd = exec.Command("ffmpeg", "-i", "bin/tmp.mp4", "-i", inputPath, "-filter_complex", "[1:v]scale=1080:607[ovr];[0:v][ovr]overlay=(main_w-overlay_w)/2:(main_h-overlay_h)/2", outputFilePath)
	fmt.Println(cmd.Args)
	
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to merge blurred and original video")
	}  

	log.Print("Successfully created video")
} 

func deleteTmpFiles() {
	err := os.RemoveAll("bin")
	if err != nil {
		log.Fatalf("Failed to delete bin folder")
	}
}
