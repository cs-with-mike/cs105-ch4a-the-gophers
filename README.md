[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/a2fQs4QM)
# Lexer & Parser for the Tokki Language
Westmont College CS 105 Fall 2023
Chapter 4, Assignments A and B

## Author Information
- **Trevor English** tenglish@westmont.edu
- **Nancy Everest** neverest@westmont.edu
- **Allie Peterson** alpeterson@westmont.edu

## Overview
### Assignment Chapter 4A
- Lexer for the Tokki language written in the Go language
- Shell Wrapper to allow the lexer to be run independent of the Go language
- Documentation for design decisions and key takeaways

### Assignment Chapter 4B
- Modify tokki.go in a new file named tokkis.go
- Add new funtions to implement the parser with recursive-descent
- Documentation for design decisions and key takeaways

### Reference
*Concepts of Programming Languages, 12th Edition by Robert Sebesta*, Chapter 4.4 Lexical and Syntax Analysis: Recursive-Descent Parsing, p 176

## Design Notes
### Assignment Chapter 4A
We designed our branching system to be based on features, to prevent merge conflicts where we're all stepping on each others toes.

Go has an int32 data type called `rune` which repesents a single Unicode character. Bytes can be easily casted to runes which allowed us to use functions from the `unicode` package (IsSpace(), IsLetter(), IsDigit()) within the getChar function to determine the charClass.

C has native getc() that ends in EOF value when the file reaches the end of the file. In our project, we use Go's `os.ReadFile()`, which doesn't give any signifier about EOF. Instead, we created a Character Class EOF with token ID (-1). Our lex() function can recongize this as the default case.

Go doesn't allow us to output the `CONST` name directly, so we created a mapping of character classes and tokens to string to let us do this.

If we just printed the lexeme casted to a string, it would contain a tail of `null` values, so we use `TrimRight()` from the `strings` package to remove the null "\x00" values.

### Assignment Chapter 4B
We were able to easily adapt the functions in the book from C to go, but we had to figure out what our initial main program would look like. We couldn't use nextToken for checking an empty file since it wouldn't be set yet to EOF, so we checked charClass instead.

We added a new global variable (depthCount) to keep track of the right and left arrows indicating the depth of the recursion tree. We also had to remove the extra spaces from the tokens that we had put in for 4A to format the output.

For the custom error message that prints the last character accessed before failure, we created a new global variable (lastChar) and saved nextChar in lastChar before changing nextChar. When it failed on nextChar, we printed that it failed at lastChar. This helps the user know where the input failed.

Once the recusrion was working, it was fairly easy to add the pretty arrow in and out formatted output.

## Lessons Learned
### Assignment Chapter 4A
We learned that Go has a relatively similar syntax (i.e. switch statements) & very similar datatypes to C, but that the functions provided in both vary. Go's vast package libraries provide some nice workarounds for the native functions used in C.

Also because auto-grader's environment (Ubuntu 22.04) did not have Go installed like it did Python, we got to learn how to write a more intricate shell script that installs Go and creates the GOCACHE env variable before running.

We also learned more about datatypes involving bytes and how they compare & contrast to that of strings.

Lastly, we learned it takes a lot of ❌ before you can get that sweet, sweet ✅ from the autograder...

### Assignment Chapter 4B
We learned how to exit out of a go program and give an error code using os.Exit(). We also learned that if you use printf to print a line using go, and if the variable is empty, it will print the % to indicate the end of line. Adding a \n to the printf removes the unwanted %.

We had to change our go.mod file each time we pushed to git. The autograder wanted go version 1.23 and locally we were running 1.21.0. This was a bit of a pain to change every time we pushed and pulled, but it was an easy workaround. Also, using the two commands (go build tokkis.go && ./tokkis $1 > out.parser.txt) rather than just go run (go run tokkis.go $1 > out.parser.txt) builds a compiled file and runs it. This eliminates the unwanted "exit 1" in the output.

We need to be careful not to push files that don't need to be changed. We had to remember to undo changes to go.mod and sample.tk before we pushed each time.

We had a few issues with people working off the same branch. Being on separate branches would not have helped in our case since we were working on the same parts of the same file. This is a difficulty that git can help with but it can't help if two people are doing the same work. Unfortunately, for some things we still have to do things the old-fashiled way--communicate manually (via Slack).