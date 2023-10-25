package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func main() {
	videoID := flag.String("id", "", "Youtube video's ID")
	outputFile := flag.String("output file name", "video.mp4", "Output filename")

	// Parse flags
	flag.Parse()
	
	if *videoID == "" {
		fmt.Println("Error: missing video ID.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	client := youtube.Client{}

	video, err := client.GetVideo(*videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create(*outputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}
