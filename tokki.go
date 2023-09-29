package main

import (
	"fmt"
	"os"
	"unicode"
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
		getChar()
	}
	fmt.Println("Contents of file:", string(contents))
}

func getNonBlank() {
	for unicode.IsSpace(nextChar) {
		getChar()
	}
}

func lex() {
	lexLen = 0

	switch charClass {
	case LETTER:
		//pass
	case DIGIT:
		//pass
	case UNKNOWN:
		//pass
	}
}

// getChar - A Function to get the next character of input and determine its character class
func getChar() {
	nextChar = contents[0]
	fmt.Printf(string(nextChar))
}

// Global Declaration Variables
var (
	charClass   int
	lexeme      [100]byte
	nextChar    rune
	lexLen      int
	token       int
	nextToken   int
	inputSource string
	contents    []rune
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
