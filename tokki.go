package main

import (
	"fmt"
	"os"
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
	} else {
		getChar(contents)

	}
	fmt.Println("Contents of file:", string(contents))
}

// getChar - A Function to get the next character of input and determine its character class
func getChar(contents []byte) {
	nextChar = []byte(contents)[0]
	fmt.Printf(string(nextChar))
}

// Global Declaration Variables
var (
	charClass   int
	lexeme      [100]byte
	nextChar    byte
	lexLen      int
	token       int
	nextToken   int
	inputSource string
)

// Character Classes
const (
	LETTER  = 0
	DIGIT   = 1
	UNKNOWN = 99
)

// Token Codes
const (
	INT_LIT     = 10
	IDENT       = 11
	ASSIGN_OP   = 20
	ADD_OP      = 21
	SUB_OP      = 22
	MULT_OP     = 23
	DIV_OP      = 24
	LEFT_PAREN  = 25
	RIGHT_PAREN = 26
)
