package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	arguments = arguments[1:]

	for _, v := range arguments {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
