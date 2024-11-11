package main

/*import "fmt"

func main() {
	fmt.Println(BalancedString("CDCCDDCDCD"))
	fmt.Println(BalancedString("CDDDDCCCDC"))
	fmt.Println(BalancedString("DDDDCCCC"))
	fmt.Println(BalancedString("CDD"))
}*/

func BalancedString(s string) int {
	cCount := 0
	dCount := 0
	balanced := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'C' {
			cCount++
		} else {
			dCount++
		}

		if cCount == dCount {
			balanced++

			cCount++
			dCount++
		}
	}
	return balanced
}
