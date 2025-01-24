package asciiart

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var inputLength int

func GenerateAscii(textToDraw, banner string) string {

	banner = "/home/nakazaki/Desktop/golang-01-nov/ascii-art-web/banners/" + banner + ".txt"

	inputLength = len(textToDraw)

	if textToDraw == "" {
		return ""
	} else if isMultipleNewLines(textToDraw) {
		return ""
	}

	file, err := os.Open(banner)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	charMapping()

	lines := strings.Split(textToDraw, "\\n")

	result := ""

	for _, line := range lines {
		if line == "" {
			fmt.Println()
			result += "\n"
		} else {
			result += printAsciiChars(banner, line)
		}
	}
	return result
}

func TextToDrawValidation(text string) error {
	for _, char := range text {
		if char < 32 || char > 126 {
			return fmt.Errorf("Your text contains some special character which is not accepted")
		}
	}
	return nil
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
