package main

import (
	"fmt"
	"log"
	"os"

	"github.build.ge.com/212472270/gocomment/input"
	"github.build.ge.com/212472270/gocomment/lint"
)

func main() {
	// Parse and handle all flags
	args := os.Args
	fileName, err := input.ParseFlags(args)
	if err != nil {
		fmt.Println("gocomment error: could not parse arguments")
		os.Exit(1)
	}

	// Get all lines for given directory
	lines, err := input.GetFileLines(fileName)
	if err != nil {
		log.Fatalln("could not get source code from " + fileName)
	}

	os.Exit(lint.Run(lines))
}
