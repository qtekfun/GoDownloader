package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

type internalData struct {
	videoID string
	outputFile string
}

func parseArgs() *internalData {
	// Define los flags
	videoID := flag.String("id", "", "Youtube video's ID")
	outputFile := flag.String("output file name", "video.mp4", "Output filename")

	// Parsea los flags desde la línea de comandos
	flag.Parse()

	// Verifica si el ID del video es proporcionado
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

func main() {
	data := parseArgs()

	client := youtube.Client{}

	video, err := client.GetVideo(data.videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create(data.outputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}
