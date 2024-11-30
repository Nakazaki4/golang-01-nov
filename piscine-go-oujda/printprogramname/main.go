package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	runes := []rune(os.Args[0])
	lastslash := 0
	for index, c := range runes {
		if c == '/' {
			lastslash = index + 1
		}
	}
	for i := lastslash; i < len(runes); i++ {
		z01.PrintRune(runes[i])
	}

	z01.PrintRune('\n')
}
