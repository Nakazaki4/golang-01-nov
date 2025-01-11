package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	arguments = arguments[1:]

	for _, c := range arguments {
		runes := []rune(c)
		for i := 0; i < len(c); i++ {
			z01.PrintRune(runes[i])
		}
		z01.PrintRune('\n')
	}
}
