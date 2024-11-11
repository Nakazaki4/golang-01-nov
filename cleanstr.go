package main

import (
	"fmt"
	"strings"
)

func CleanStr(str string) string {
	if str == "" {
		return ""
	}
	strings.ReplaceAll()
}

func main() {
	fmt.Println(CleanStr("you see it's easy to display the same thing"))
	fmt.Println(CleanStr(" only  it's herder  "))
	fmt.Println(CleanStr(""))
}
