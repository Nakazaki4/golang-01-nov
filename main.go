package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	fmt.Println(Fprime(225225))
}

func Fprime(nb int) string {
	if nb < 2{
		return ""
	}
	primes := []int{}
	for i := 2; i*i <= nb; i++ {
		for nb%i == 0 {
			primes = append(primes, i)
			nb /= i
		}
	}

	
	if nb > 1 {
		primes = append(primes, nb)
	}

	str := ""
	for i, v := range primes {
		if i > 0 {
			str += "*"
		}
		temp := ""
		for v > 0 {
			temp = string('0'+v%10) + temp
			v /= 10
		}
		str += temp
	}

	return str
}

func ConcatAlternate(slice1, slice2 []int) []int {
	if len(slice1) == 0 {
		return slice2
	}
	if len(slice2) == 0 {
		return slice1
	}

	res := []int{}

	for i, v := range slice1 {
		if len(slice1) > len(slice2) {
			res = append(res, v, slice2[i])
		} else if len(slice2) > len(slice1) {
			res = append(res, slice2[i], v)
		} else {
			res = append(res, v, slice2[i])
		}
	}
	if len(slice1) > len(slice2) {
		res = append(res, slice1[len(slice2):]...)
	} else if len(slice2) > len(slice1) {
		res = append(res, slice2[len(slice1):]...)
	}

	return res
}

func AddPrimeSum(nb int) int {
	res := 0

	for i := 2; i <= nb; i++ {
		if isPrime(i) {
			res += i
		}
	}
	return res
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Print()
	} else {

		res := [][]int{}
		for i := 0; i < len(slice); i += size {
			end := i + size
			if end > len(slice) {
				end = len(slice)
			}
			res = append(res, slice[i:end])
		}
		fmt.Print(res)
	}
	fmt.Println()
}

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	uniqueChars1 := make(map[rune]bool)
	uniqueChars2 := make(map[rune]bool)

	for _, v := range str1 {
		uniqueChars1[v] = true
	}
	for _, v := range str2 {
		uniqueChars2[v] = true
	}

	count := 0

	for char := range uniqueChars1 {
		if !uniqueChars2[char] {
			count++
		}
	}

	for char := range uniqueChars2 {
		if !uniqueChars1[char] {
			count++
		}
	}
	return count
}

func ThirdTimeIsACharm(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 2 {
		return "\n"
	}
	size := len(str)
	res := ""

	for i := 2; i < size; i += 3 {
		res += string(str[i])
	}
	return res + "\n"
}

func Printrevcomb() {
	for i := 9; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			for k := j - 1; k >= 0; k-- {
				z01.PrintRune(rune(i + '0'))
				z01.PrintRune(rune(j + '0'))
				z01.PrintRune(rune(k + '0'))
				if !(i == 2 && j == 1 && k == 0) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
