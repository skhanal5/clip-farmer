package download

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/skhanal5/clip-farmer/internal/client"
)

const (
	connectTimeout = 30 * time.Second
	chunkSize      = 1024 * 1024
)

// downloadMP4File handles the logic to download a clip given an url, the name of the clip, and a path
// to write the contents of the clip to.
func DownloadMP4File(mp4Link string, outputPath string) {
	responseBody := getMP4URL(mp4Link)
	writeBodyAsMP4(responseBody, outputPath)
}

// getMP4URL handles sending a GET request to the URL containing
// the raw mp4 file. 
// Returns the response body of the GET request.
func getMP4URL(url string) io.ReadCloser {
	log.Printf("Getting MP4 file at URL: %s\n", url)
	responseBody, err := client.GetURL(url)
	if err != nil {
		log.Fatalf("Received error: %v after attempting to retrieve MP4 file\n", err)
	}
	return responseBody
}

// writeBodyAsMP4 writes the response body as
// a mp4 file to the output path
func writeBodyAsMP4(responseBody io.ReadCloser, outputPath string) {
	log.Printf("Writing MP4 file: %s\n", outputPath)
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Failed to create output file path with error: %v\n", err)
	}
	writeChunksToFile(responseBody, file)
}

// writeChunksToFile will write the contents of the
// get request chunk by chunk onto the file
// Referenced from twitch-dll
func writeChunksToFile(responseBody io.ReadCloser, file *os.File) {
	size := int64(0)
	buf := make([]byte, chunkSize)
	for {
		n, err := responseBody.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}

		_, err = file.Write(buf[:n])
		if err != nil {
			log.Fatal(err)
		}
		size += int64(n)
	}
}