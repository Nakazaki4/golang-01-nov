package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var textToDrawLen int

func main() {
	args := os.Args

	var textToDraw string
	bannerToUse := "/home/aboudchar/Desktop/golang-01-nov/ascii-art/standard.txt"

	if len(args) != 2 {
		log.Fatal("You should input exactly one argument")
	}

	textToDraw = args[1]
	textToDrawLen = len(textToDraw)
	if textToDraw == "" {
		return
	} else if textToDraw == "\\n" {
		fmt.Println()
		return
	} else if multipleNewLines(textToDraw) {
		for i := 0; i < textToDrawLen/2; i++ {
			fmt.Println()
		}
		return
	}
	textToDrawValidation(textToDraw)

	file, err := os.Open(bannerToUse)
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
			printAsciiChars(bannerToUse, line)
		}
	}
}

func textToDrawValidation(text string) {
	for _, char := range text {
		if char < 32 || char > 126 {
			log.Fatal("Your text contains a special character")
		}
	}
}

func multipleNewLines(text string) bool {
	for i := 0; i < textToDrawLen; i++ {
		if i < textToDrawLen-1 && !(text[i] == '\\' && text[i+1] == 'n') {
			return false
		} else {
			i++
		}
	}
	return true
}
