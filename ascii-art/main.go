package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args

	var textToDraw string
	bannerToUse := "/home/aboudchar/Desktop/golang-01-nov/ascii-art/standard.txt"

	if len(args) != 2 {
		log.Fatal("You should input exactly one argument")
	}

	textToDraw = args[1]
	if textToDraw == "" {
		return
	}else if textToDraw == "\\n"{
		fmt.Println()
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
