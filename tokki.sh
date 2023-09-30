#!/bin/bash
# Authors:
#  - Trevor English, tenglish@westmont.edu
#  - Nancy Everest, neverest@westmont.edu
#  - Allie Peterson, alpeterson@westmont.edu

# Run any commands necessary to set up your language's runtime environment here.
# If the runtime is expected to be present on Ubuntu by default, then do nothing.
echo "Setting up runtime ..."

# Install Go. (mryu)
sudo apt install golang; go version

# Setup GOCACHE environment variable. (mryu)
echo "Creating directory $(pwd)/go_cache and setting it to GOCACHE envvar."
mkdir -p ./go_cache; export GOCACHE=$(pwd)/go_cache

# Run your Tokki lexer, passing in the first command line argument directly to the lexer.
echo "Running Tokki ..."

# Here is the command to run & output it.
go run tokki.go $1 > out.txt