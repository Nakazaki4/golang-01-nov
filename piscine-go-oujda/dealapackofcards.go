package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(d []int) {
	a := 4
	b := 3

	for c := 1; c <= a; c++ {
		fmt.Printf("Player %d: ", c)

		for e := 0; e < b; e++ {
			fmt.Printf("%d", d[(c-1)*b+e])
			if e < b-1 {
				fmt.Printf(", ")
			}
		}

		z01.PrintRune('\n')
	}
}
