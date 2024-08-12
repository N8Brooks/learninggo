package main

import (
	"embed"
	"fmt"
	"os"
)

//go:embed human_rights
var humanRights embed.FS

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a language")
		os.Exit(1)
	}
	data, err := humanRights.ReadFile("human_rights/" + os.Args[1] + ".txt")
	if err != nil {
		fmt.Println("Language not found")
		os.Exit(1)
	}
	fmt.Println(string(data))
}
