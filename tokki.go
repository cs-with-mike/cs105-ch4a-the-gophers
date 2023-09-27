package main

import (
	"fmt"
	"os"
)

// main function, reads filename from second argument
// usage: go run . sample.tk
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(contents))
}
