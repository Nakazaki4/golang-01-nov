package piscine

func BasicAtoi2(s string) int {
	var num int
	for a := 0; a < len(s); a++ {
		res := int(s[a] - '0')
		if res < 0 || res > 9 {
			return 0
		}
		num = num*10 + res
	}
	return num
}
