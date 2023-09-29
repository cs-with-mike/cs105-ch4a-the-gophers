#!/bin/bash
# Authors:
#  - Boaty McBoatface, bmcboatface@westmont.edu
#  - Thomas McTankengine, tmctankengine@westmont.edu
#  - Ovenly McEasybake, omceasybake@westmont.edu

# Run any commands necessary to set up your language's runtime environment here.
# If the runtime is expected to be present on Ubuntu by default, then do nothing.
echo "Setting up runtime ..."

echo "Installing Homebrew"

ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

echo "Installing Go..."

brew update&& brew install golang

# I'm checking Go interpreter's version here just as a placeholder.
go version # TODO: YOUR RUNTIME SETUP HERE.

# Run your Tokki lexer, passing in the first command line argument directly to the lexer.
# Any output to STDOUT should be directed to a text output file titled "out.txt."
echo "Running Tokki ..."

# As an example, I have provided how I would run my tokki.pyc.
go run tokki.go $1 > out.txt # TODO: REPLACE THIS WITH YOUR OWN COMMAND.