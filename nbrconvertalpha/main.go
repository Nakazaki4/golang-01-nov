package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	deci := 96

	if arguments[1] == "--upper" {
		deci = 64
	}

	for _, c := range arguments[2:] {
		num := 0
		for i := 0; i < len(c); i++ {
			num = num*10 + int(c[i]-'0')
		}
		if num >= 1 && num <= 27 {
			z01.PrintRune(rune(num + deci))
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
