package main

import (
	"fmt"
	"bufio"
	"os"
)

var charMap = make(map[rune]int)

func charMapping() {
	chars := []rune{
		'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		':', ';', '<', '=', '>', '?', '@',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'[', '\\', ']', '^', '_', '`',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'{', '|', '}', '~',
	}

	index := 12
	for i := 0; i < len(chars); i++ {
		charMap[chars[i]] = index
		index += 9
	}
}

func getCharAddress(text string) []int {
	runes := []rune(text)
	addr := make([]int, len(text))

	for i, r := range runes {
		addr[i] = charMap[r]
	}
	return addr
}

func getCharsAscii(banner string, lines []int) error {
	content, err := os.Open(banner)
	if err != nil {
		return err
	}
	defer content.Close()

	r := bufio.NewScanner(content)

	var allLines []string

	for r.Scan() {
		allLines = append(allLines, r.Text())
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < len(lines); j++ {
			start := lines[j] - 1
			if start+i < len(allLines) {
				c := allLines[start+i]
				fmt.Print(c)
			}
		}
		fmt.Println()
	}
	return err
}
