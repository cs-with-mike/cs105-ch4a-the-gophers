# Lexer & Parser for the Tokki Language
Westmont College CS 105 Fall 2023
Chapter 4 Assignment A

## Author Information
- **Trevor English** tenglish@westmont.edu
- **Nancy Everest** neverest@westmont.edu
- **Allie Peterson** alpeterson@westmont.edu

## Overview
Assignment Chapter 4A
- Lexer for the Tokki language written in the Go language
- Shell Wrapper to allow the lexer to be run independent of the Go language
- Documentation for design decisions and key takeaways

## Design Notes
Go has an int32 data type called `rune` which repesents a single Unicode character. Bytes can be easily casted to runes which allowed us to use functions from the `unicode` package (IsSpace(), IsLetter(), IsDigit())

C has native getc() that ends in EOF value when the file reaches the end of the file. In our project, we use Go's os.ReadFile() instead...
## Lessons Learned
(TODO: Describe any lessons learned while you were working on this assignment.)
