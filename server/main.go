package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir := "./audios"
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Print(file)
	}
}
