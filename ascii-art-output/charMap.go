package main

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

func printAsciiChars(text string, allLines []string) string {
	result := ""
	for i := 0; i < 8; i++ {
		for _, char := range text {
			start := getCharAddress(char) - 1
			if start+i < len(allLines) && start+i >= 0 {
				result += allLines[start+i]
			}
		}
		result += "\n"
	}
	return result
}
