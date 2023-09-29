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
	if err != nil {
		fmt.Println("File reading error", err)
		return
	} else {
		getChar()
		for nextToken != EOF {
			lex()
		}

	}
	//fmt.Println("Contents of file:", string(contents))
}

// Lookup - a function to lookup operators and parentheses and return the token
func lookup(char byte) int {
	switch char {
	case '(':
		addChar()
		nextToken = LEFT_PAREN
		break
	case ')':
		addChar()
		nextToken = RIGHT_PAREN
		break
	case '+':
		addChar()
		nextToken = ADD_OP
		break
	case '-':
		addChar()
		nextToken = SUB_OP
		break
	case '*':
		addChar()
		nextToken = MULT_OP
		break
	case '/':
		addChar()
		nextToken = DIV_OP
		break
	default:
		addChar()
		nextToken = EOF
		break
	}
	return nextToken
}

func getNonBlank() {
	for unicode.IsSpace(rune(nextChar)) {
		getChar()
	}
}

func lex() int {
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

	// Parentheses and operators
	case UNKNOWN:
		lookup(nextChar)
		getChar()
		break

	// EOF
	case EOF:
		nextToken = EOF
		lexeme[0] = 'E'
		lexeme[1] = 'O'
		lexeme[2] = 'F'
		lexeme[3] = 0
		lexeme[4] = 0
		break
	} // End of switch
	fmt.Printf("Next token is: %s | Next lexeme is %s\n", tokensCharClassesToString[nextToken], lexeme)
	return nextToken
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
	if len(contents) > 0 {
		nextChar = contents[0]
		contents = contents[1:]
		if nextChar != 0 {
			if unicode.IsLetter(rune(nextChar)) {
				charClass = LETTER
			} else if unicode.IsDigit(rune(nextChar)) {
				charClass = DIGIT
			} else {
				charClass = UNKNOWN
			}
		} else {
			charClass = EOF
		}
	} else {
		charClass = EOF
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
	EOF     = -1
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

// Map tokens and character classes to string for outputting
var tokensCharClassesToString = map[int]string{
	INT_LIT:     "INT_LIT",
	IDENT:       "IDENT",
	ASSIGN_OP:   "ASSIGN_OP",
	ADD_OP:      "ADD_OP",
	SUB_OP:      "SUB_OP",
	MULT_OP:     "MULT_OP",
	DIV_OP:      "DIV_OP",
	LEFT_PAREN:  "LEFT_PAREN",
	RIGHT_PAREN: "RIGHT_PAREN",
	LETTER:      "LETTER",
	DIGIT:       "DIGIT",
	UNKNOWN:     "UNKNOWN",
	EOF:         "EOF",
}
