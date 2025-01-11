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

	
	arr := []string{"h","e","l","l","o"}
	j := len(arr)-1
	for i:=0; i<len(arr)/2; i++ {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
		j--
	}
	fmt.Println(arr)
}
