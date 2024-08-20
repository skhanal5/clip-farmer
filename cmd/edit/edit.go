package edit

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	inputPath string
	outputPath string
	blurredOption bool
)

// configCmd represents the config command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a video",
	RunE: func(cmd *cobra.Command, args []string) error {
		if blurredOption != false {
			CreateVideoWithBlurredBackground(inputPath, outputPath)
		}		
		return nil
	},
}

func Init() *cobra.Command {
	editCmd.Flags().StringVarP(&inputPath, "input", "i", "",
		"Path of the mp4 file that we would like to edit")
	editCmd.Flags().StringVarP(&outputPath, "output", "o", "",
		"Path of the resulting edited video.")
	editCmd.MarkFlagsRequiredTogether("input", "output")
	
	
	editCmd.Flags().BoolVarP(&blurredOption, "blurred", "b", false,
	"Make a video that is overlayed ontop of a blurred background")
	return editCmd
}

func CreateVideoWithBlurredBackground(input string, output string) {
	log.Print("Blurring the video")
	
	// Change resolution to TikTok's viewport
	// Apply box blur
	// Remove audio
	cmd := exec.Command("ffmpeg", "-i", input, "-s", "1080x1920", "-vf", "boxblur=50", "-an", "tmp.mp4")
	fmt.Println(cmd.Args)
	_, err := cmd.CombinedOutput()
	
	// delete video using defer
	defer deleteTmpFiles()
	
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Failed to blur video")
	} 
	
	// then overlay the video with the "clear" version
	log.Print("Overlaying the original video on top of the blurred video.")
	cmd = exec.Command("ffmpeg", "-i", "tmp.mp4", "-i", input, "-filter_complex", "[1:v]scale=1080:607[ovr];[0:v][ovr]overlay=(main_w-overlay_w)/2:(main_h-overlay_h)/2", output)
	fmt.Println(cmd.Args)
	
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Failed to merge blurred and original video")
	}  

	log.Print("Successfully created video")
} 

func deleteTmpFiles() {
	err := os.Remove("tmp.mp4")
	if err != nil {
		log.Fatalf("Failed to delete file")
	}
}


// func CreateVideoWithStackedBackground() {

// }
