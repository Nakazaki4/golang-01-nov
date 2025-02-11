package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	inputLength    int
	textToDraw     string
	outputFilename string
	resultString   string
)

func main() {
	arguments := os.Args

	if len(arguments) != 4 {
		log.Fatal("Usage: go run . [--output=<fileName.txt>] <STRING> [BANNER]")
	}

	textToDraw = arguments[2]
	outputFilename = arguments[1]
	bannerFile := arguments[3]

	if !fnValidation(outputFilename) {
		log.Fatal("Usage: go run . [--output=<fileName.txt>] <STRING> [BANNER]")
	}

	bannerFilePath := bannerFile + ".txt"

	inputLength = len(textToDraw)

	if textToDraw == "" {
		return
	} else if textToDraw == "\n" {
		resultString += "\n"
	} else if isMultipleNewLines(textToDraw) {
		for i := 0; i < inputLength/2; i++ {
			resultString += "\n"
		}
		WriteFile(resultString)
		return
	}

	textToDrawValidation(textToDraw)
	allLines := bannerScanner(bannerFilePath)

	charMapping()
	if textContainsN(textToDraw) {
		lines := strings.Split(textToDraw, "\\n")

		for _, line := range lines {
			if line == "" {
				resultString += "\n"
			} else {
				resultString += printAsciiChars(line, allLines)
			}
		}
	} else {
		resultString += printAsciiChars(textToDraw, allLines)
	}
	WriteFile(resultString)
}

func textContainsN(text string) bool {
	for i := 0; i < len(text); i++ {
		if i < inputLength-1 && !(text[i] == '\\' && text[i+1] == 'n') {
			return true
		}
	}
	return false
}

func fnValidation(arg string) bool {
	r := regexp.MustCompile(`--output=.*\.txt`)
	return r.MatchString(arg)
}

func textToDrawValidation(text string) {
	for _, char := range text {
		if char < 32 || char > 126 {
			log.Fatal("Your text contains some special character which is not accepted.")
		}
	}
}

func isMultipleNewLines(text string) bool {
	if text == "\\" {
		return false
	}
	for i := 0; i < inputLength; i += 2 {
		if !(text[i] == '\\' && text[i+1] == 'n') {
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

	for r.Scan() {
		allLines = append(allLines, r.Text())
	}

	return allLines
}

func WriteFile(text string) {
	outputFilename = strings.ReplaceAll(outputFilename, "--output=", "")
	if outputFilename == "standard.txt" || outputFilename == "shadow.txt" || outputFilename == "thinkertoy.txt" {
		log.Fatal("Choose another name for the ouput file.")
	}
	os.WriteFile(outputFilename, []byte(text), 0o644)
}
