package main

import (
	"os"
)

func main() {
	audioFilesPath := "./audios"
	os.Mkdir("processedAudios", 0755)
	files, err := os.ReadDir(audioFilesPath)
	if err != nil {
		panic("Something went wrong")
	}
	for _, file := range files {

	}
}
