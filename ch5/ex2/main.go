package main

import (
	"log"
	"os"
)

func fileLen(name string) (int64, error) {
	file, err := os.Open(name)
	if err != nil {
		return -1, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return -1, err
	}
	return fileInfo.Size(), nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main <file>")
	}
	count, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	println(count)
}
