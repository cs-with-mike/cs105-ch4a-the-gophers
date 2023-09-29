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
		fmt.Println("Missing parameter, provide input file name!")
		return
	}
	var err error
	contents, err = os.ReadFile(os.Args[1])
	fmt.Println(contents)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	} else {
		getChar()
		// for nextToken != EOF {
		// 	lex()
		// }

	}
	fmt.Println("Contents of file:", string(contents))
}

func getNonBlank() {
	for unicode.IsSpace(rune(nextChar)) {
		getChar()
	}
}

func lex() {
	lexLen = 0
	getNonBlank()

	switch charClass {
	// Parse identifiers
	case LETTER:
		addChar()
		getChar()
		for charClass == LETTER || charClass == DIGIT {
			addChar()
			getChar()
		}
		nextToken = IDENT
		break

	// Parse integer literals
	case DIGIT:
		addChar()
		getChar()
		for charClass == DIGIT {
			addChar()
			getChar()
		}
		nextToken = INT_LIT
		break

	case UNKNOWN:
		//lookup(nextChar)
		getChar()
		break
	}
	//case EOF
	//pass
}

// addChar - a function to add nextChar to lexeme
func addChar() {
	if lexLen <= 98 {
		lexeme[lexLen] = nextChar
		lexLen++
		lexeme[lexLen] = 0
	} else {
		fmt.Println("Error - lexeme is too long")
	}
}

// getChar - A Function to get the next character of input and determine its character class
func getChar() {
	nextChar = contents[0]
	if unicode.IsLetter(rune(nextChar)) {
		charClass = LETTER
	} else if unicode.IsDigit(rune(nextChar)) {
		charClass = DIGIT
	} else {
		charClass = UNKNOWN
	}

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
	contents    []byte
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
