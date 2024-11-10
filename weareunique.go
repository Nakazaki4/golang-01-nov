package main

import "fmt"

func main() {

	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(str1, str2 string) int {
	if str1 =="" && str2 == "" {
		return -1
	}

	count := 0

	for i:= 0; i <= len(str1)-1; i++{
		for j := 1; j <= len(str2)-1; j++ {
			if !(str1[j] == str2[i]) {
				count++
			}
		}
	}
	return count
}
