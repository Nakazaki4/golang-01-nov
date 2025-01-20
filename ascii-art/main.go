package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var inputLength int

func main() {
	args := os.Args

	bannerFilePath := "/home/aboudchar/Desktop/golang-01-nov/ascii-art/standard.txt"

	if len(args) != 2 {
		log.Fatal("You should input exactly one argument")
	}

	textToDraw := args[1]
	inputLength = len(textToDraw)

	if textToDraw == "" {
		return
	} else if textToDraw == "\\n" {
		fmt.Println()
		return
	} else if isMultipleNewLines(textToDraw) {
		for i := 0; i < inputLength/2; i++ {
			fmt.Println()
		}
		return
	}

	textToDrawValidation(textToDraw)

	file, err := os.Open(bannerFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	charMapping()

	lines := strings.Split(textToDraw, "\\n")

	for _, line := range lines {
		if line == "" {
			fmt.Println()
		} else {
			printAsciiChars(bannerFilePath, line)
		}
	}
}

func textToDrawValidation(text string) {
	for _, char := range text {
		if char < 32 || char > 126 {
			log.Fatal("Your text contains some special character which is not accepted")
		}
	}
}

func isMultipleNewLines(text string) bool {
	for i := 0; i < inputLength; i++ {
		if i < inputLength-1 && !(text[i] == '\\' && text[i+1] == 'n') {
			return false
		} else {
			i++
		}
	}
	return true
}
