package main

// eof menas end of file

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// main function, reads filename from second argument
// usage: go run tokki.go sample.tk
func main() {
	// if there arnt 2 arguments when you try to run this
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide input file name!")
		return
	}
	var err error // declares a variable lol

	// sets contents of the file
	// puls everything from first file into contents and put errors into this error var
	contents, err = os.ReadFile(os.Args[1])
	//if its not clean-- no errors , so it can move to else
	if err != nil {
		fmt.Println("File reading error", err)
		// return stops it
		return
	} else {
		// now its clean and we can call get car

		getChar()
		// starts blank, is an int
		// while the next token isnt the end, call the lexical
		for nextToken != EOF {
			lex()
		}

	}
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

func lex() int { // called every time we hit a new lexeme
	// length of lexeme starts at 0, we dont have lexeme yet
	lexLen = 0
	// separating at the blank spaces-- this is why we can ignore spaces
	getNonBlank()
	// switch statement similar to conditionals in racket-- same as if
	switch charClass {
	// Parse identifiers
	// is the case is a letter:
	case LETTER:
		addChar() // adds the char to the lexeme
		getChar() // gets the next one (parces it)

		// until the next one is isnt a letter-- this is how we finish words.
		// digits okay bs of digits in var names
		for charClass == LETTER || charClass == DIGIT {
			addChar()
			getChar()
		}
		// ident is a string, sets the type of the lexeme to ident which is string here
		nextToken = IDENT
		break

	// Parse integer literals
	// basuically same as string one
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
		// lookup calls the function and finds the opperator that the unknown is
		lookup(nextChar)
		getChar() //sets the next so stars over)

		break

	// EOF
	//if eof, prints this
	case EOF:
		nextToken = EOF
		lexeme[0] = 'E'
		lexeme[1] = 'O'
		lexeme[2] = 'F'
		// we only want it to print EOF, not EOF puss 97 null spaces
		// everry leveme after F is 0, so its trimmed
		for i := 3; i < len(lexeme); i++ {
			lexeme[i] = 0
		}
		break
	} // End of switch
	fmt.Printf("Next token is: %s | Next lexeme is %s\n", tokensCharClassesToString[nextToken], strings.TrimRight(string(lexeme[:]), "\x00"))
	return nextToken // next token will be what the case defined it as (int_lit) or whatever is being printed
}

// addChar - a function to add nextChar to lexeme
// called in lookup and in our lex function
// actually adding what we have to the lexeme
func addChar() {
	// starts at 0, so length is only within the call of lex
	// 98 bc?
	if lexLen <= 98 {
		// sets the first inex to the next char and increment the counter up by one
		lexeme[lexLen] = nextChar
		lexLen++
		// a formatting thing-- for every char that isnt eof. make the rest of it 0 values, not null. zeroes get deleted when we print
		for i := lexLen; i < 100; i++ {
			lexeme[i] = 0
		}
	} else {
		fmt.Println("Error - lexeme is too long") // if we had a lexeme more than 100 digits lol
	}
}

// getChar - A Function to get the next character of input and determine its character class
func getChar() {
	// if it has anything in it
	if len(contents) > 0 {

		// basically "get c" function for go
		// were focusing at next chat, which is index 0
		nextChar = contents[0]
		// poping cotents -- right away we move contents forward
		contents = contents[1:]

		if nextChar != 0 { // if there IS a next char
			// sets a char class
			// rune is unicode rep of byte (bc ascii has less divers char library and uses bytes)
			if unicode.IsLetter(rune(nextChar)) {
				// settign chaar class ("go down this route")
				charClass = LETTER
			} else if unicode.IsDigit(rune(nextChar)) {
				/// settign chaar class ("go down this route")
				charClass = DIGIT
			} else {
				// settign chaar class ("go down this route")
				charClass = UNKNOWN
			}
		} else { // we reach the end. 0 is null char
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
// nums just make unique
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
	INT_LIT:     "    INT_LIT",
	IDENT:       "      IDENT",
	ASSIGN_OP:   "  ASSIGN_OP",
	ADD_OP:      "     ADD_OP",
	SUB_OP:      "     SUB_OP",
	MULT_OP:     "    MULT_OP",
	DIV_OP:      "     DIV_OP",
	LEFT_PAREN:  " LEFT_PAREN",
	RIGHT_PAREN: "RIGHT_PAREN",
	LETTER:      "     LETTER",
	DIGIT:       "      DIGIT",
	UNKNOWN:     "    UNKNOWN",
	EOF:         "        EOF",
}
