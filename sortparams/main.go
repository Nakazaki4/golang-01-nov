package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	arguments = arguments[1:]
	runes := []rune{}

	sort(&arguments)

	for _, c := range arguments {
		for _, j := range c {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}

	for _, x := range runes {
		z01.PrintRune(x)
		z01.PrintRune('\n')
	}
}

func sort(slice *[]string) {
	for i := 0; i < len(*slice); i++ {
		for j := i + 1; j < len(*slice); j++ {
			if (*slice)[j] < (*slice)[i] {
				(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
			}
		}
	}
}
