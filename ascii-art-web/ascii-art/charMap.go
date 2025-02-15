package asciiart

import (
	"bufio"
	"log"
	"os"
)

var charMap = make(map[rune]int)

func charMapping() {
	chars := []rune{
		' ', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		':', ';', '<', '=', '>', '?', '@',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'[', '\\', ']', '^', '_', '`',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'{', '|', '}', '~',
	}

	index := 3
	for i := 0; i < len(chars); i++ {
		charMap[chars[i]] = index
		index += 9
	}
}

func getCharAddress(r rune) int {
	return charMap[r]
}

func printAsciiChars(banner string, text string) string {
	// Open the banner file
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

	result := ""

	// Convert each character in `text` to its corresponding ASCII lines
	for i := 0; i < 8; i++ { // 8 lines per character in the ASCII art
		for _, char := range text {
			start := getCharAddress(char) - 1
			if start+i < len(allLines) && start+i >= 0 {
				result += allLines[start+i]
			}
		}
		result += "\n" // Move to the next line of ASCII art
	}
	return result
}
