package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args

	var outputFilename string
	//var textToDraw string
	var bannerToUse string

	// Cheking for the missing required arguments
	if len(args) < 2 {
		log.Fatal("You forgot to give a name to the output file.")
	} else if len(args) < 3 {
		log.Fatal("You forgot to add the text to draw.")
	} else if len(args) < 4 {
		log.Fatal("You didn't specify the banner.")
	}

	// Validating the input file name
	if strings.HasPrefix(args[1], "--output=") {
		outputFilename = args[1]
		r := regexp.MustCompile(`--output=(\w+.txt)`)
		outputFilename = r.FindString(outputFilename)
		outputFilename = r.ReplaceAllString(outputFilename, "$1")
		if !(strings.HasSuffix(outputFilename, ".txt")) {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(1)
		}
	} else {
		log.Fatal("Invalid output file name")
	}

	//text to draw = args[2]
	bannerToUse = args[3]

	// Checking if the banner file exists
	_, error := os.Open(bannerToUse + ".txt")

	if error != nil {
		log.Fatal(error)
	}

}
