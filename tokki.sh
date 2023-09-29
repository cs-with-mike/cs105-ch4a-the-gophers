#!/bin/bash
# Authors:
#  - Trevor English, tenglish@westmont.edu
#  - Nancy Everest, neverest@westmont.edu
#  - Allie Peterson, alpeterson@westmont.edu

# Run any commands necessary to set up your language's runtime environment here.
# If the runtime is expected to be present on Ubuntu by default, then do nothing.
echo "Setting up runtime ..."

# I'm checking Go interpreter's version here just as a placeholder.
go version # TODO: YOUR RUNTIME SETUP HERE.

chmod 744 tokki.go

# Run your Tokki lexer, passing in the first command line argument directly to the lexer.
# Any output to STDOUT should be directed to a text output file titled "out.txt."
echo "Running Tokki ..."

# As an example, I have provided how I would run my tokki.pyc.
./tokki.go sample.tk > out.txt