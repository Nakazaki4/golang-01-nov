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

	var outputFilename, textToDraw, bannerToUse string

	if len(args) < 2 || len(args) < 3 || len(args) < 4 {
		errorMessage()
	}

	if strings.HasPrefix(args[1], "--output=") {
		outputFilename = args[1]
		r := regexp.MustCompile(`--output=(\w+.txt)$`)
		outputFilename = r.ReplaceAllString(r.FindString(outputFilename), "$1")
		if !(strings.HasSuffix(outputFilename, ".txt")) {
			errorMessage()
		}
	} else {
		log.Fatal("Invalid output file name")
	}

	textToDraw = args[2]
	textToDrawValidation(textToDraw)

	bannerToUse = args[3]

	file, error := os.Open(bannerToUse + ".txt")
	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()

	charMapping()
	err := getCharsAscii(bannerToUse+".txt", getCharAddress(textToDraw))
	if err != nil {
		log.Fatal(err)
	}
}

func textToDrawValidation(text string) {
	runes := []rune(text)

	for _, char := range runes {
		if char < 32 || char > 126 {
			log.Fatal("Your text contains a special character")
		}
	}
}

// Prints the formating instructions and Exits with code 1
func errorMessage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> something standard")
	os.Exit(1)
}
