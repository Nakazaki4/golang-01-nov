package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var inputLength int

func main() {
	args := os.Args

	bannerFilePath := "/home/nakazaki/Desktop/golang-01-nov/ascii-art/standard.txt"

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
	allLines := bannerScanner(bannerFilePath)

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
			printAsciiChars(line, allLines)
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
	if text == "\\" {
		return false
	}
	for i := 0; i < inputLength; i++ {
		if i < inputLength-1 && !(text[i] == '\\' && text[i+1] == 'n') {
			return false
		}
	}
	return true
}

func bannerScanner(banner string) []string {
	content, err := os.Open(banner)
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()

	r := bufio.NewScanner(content)
	var allLines []string

	// Read the entire banner into memory
	for r.Scan() {
		allLines = append(allLines, r.Text())
	}
	return allLines
}
