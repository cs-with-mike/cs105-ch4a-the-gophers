package main

import (
	"fmt"
	"os"
	"io/ioutil" // not sure if we need this, but we might for input and output patterns 
)

// main function, reads filename from second argument
// usage: go run tokki.go sample.tk
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

// Global Declarations 
var (
	charClass int 
	lexeme [100]byte
	nextChar byte
	lexLen int 
	token int
	nextToken int 
	inputSource string 
)
