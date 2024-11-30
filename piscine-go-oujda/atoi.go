package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	var num int
	isNegative := false

	if s[0] == '-' {
		isNegative = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for a := 0; a < len(s); a++ {
		res := int(s[a] - '0')
		if res < 0 || res > 9 {
			return 0
		}
		num = num*10 + res
	}

	if isNegative {
		num = -num
	}

	return num
}
