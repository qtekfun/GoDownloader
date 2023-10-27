package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2"
)

type internalData struct {
	videoID    string
	outputFile string
}

func parseArgs() *internalData {
	// Define args
	videoID := flag.String("id", "", "Youtube video's ID")
	outputFile := flag.String("output file name", "audio.mp4", "Output filename")

	// Parse args
	flag.Parse()

	// VCheck if some video is requested
	if *videoID == "" {
		fmt.Println("Error: missing video ID.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	return &internalData{
		videoID:    *videoID,
		outputFile: *outputFile,
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data := parseArgs()

	client := youtube.Client{}

	video, err := client.GetVideo(data.videoID)
	checkError(err)

	var selectedFormat *youtube.Format
	// Encuentra el formato de audio MP3
	for _, format := range video.Formats.WithAudioChannels() {
		if strings.Contains(format.MimeType, "audio/mp4") {
			selectedFormat = &format
			break
		}
	}

	if selectedFormat == nil {
		fmt.Println("Error: No audio mp4 found for this ID.")
		os.Exit(1)
	}

	stream, _, err := client.GetStream(video, selectedFormat)
	checkError(err)
	defer stream.Close()

	file, err := os.Create(data.outputFile)
	checkError(err)
	defer file.Close()

	_, err = io.Copy(file, stream)
	checkError(err)
}
