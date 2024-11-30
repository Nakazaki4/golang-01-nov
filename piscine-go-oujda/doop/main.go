package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	intNumber1 := 0
	intNumber2 := 0

	if len(os.Args) != 4 {
		return
	}

	if isNumber(args[1]) || isNumber(args[3]) {
		fmt.Println("OHO")
		return
	} else {
		intNumber1 = toInt(args[1])
		intNumber2 = toInt(args[3])
	}

	operator := args[2]
	if operator == "+" {
		res, error := addAndSub(intNumber1, intNumber2)
		if !error {
			return
		}
		fmt.Println(res)

	} else if operator == "-" {
		res, error := addAndSub(intNumber1, intNumber2)
		if !error {
			return
		}
		fmt.Println(res)

	} else if operator == "*" {
		res, error := multiply(intNumber1, intNumber2)
		if !error {
			return
		}
		fmt.Println(res)

	} else if operator == "/" {
		res, error := divideAndModulo(intNumber1, intNumber2)
		if !error {
			fmt.Println("No division by 0")
		}
		fmt.Println(res)

	} else if operator == "%" {
		res, error := divideAndModulo(intNumber1, intNumber2)
		if !error {
			fmt.Println("No modulo by 0")
		}
		fmt.Println(res)
	}
}

func isNumber(n string) bool {
	for _, v := range n {
		if !(v >= '0' && v <= '9') {
			return true
		}
	}
	return false
}

func toInt(n string) int {
	number := 0
	sign := 1
	fdigi := 0

	if n[0] == '-' {
		sign = -1
	}
	for i := fdigi; i < len(n); i++ {
		if n[i] >= '0' && n[i] <= '9' {
			number = number*10 + int(n[i]-'0')
		}
	}
	return number * sign
}

func addAndSub(a, b int) (int, bool) {
	sum := a + b
	if (a > 0 && b > 0 && sum < 0) || (a < 0 && b < 0 && sum > 0) {
		return 0, false
	}
	return sum, true
}

func divideAndModulo(a, b int) (int, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}
	res := a / b
	return res, true
}

func multiply(a, b int) (int, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}

	if a > 9223372036854775807/b || a < -9223372036854775807/b {
		return 0, false
	}
	return a * b, true
}
