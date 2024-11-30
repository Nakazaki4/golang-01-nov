package piscine

func BasicAtoi(s string) int {
	runes := []rune(s)

	var num int

	for a := 0; a < len(runes); a++ {
		for b := '0'; b <= '9'; b++ {
			if runes[a] == b {
				num = num*10 + int(b-'0')
			}
		}
	}
	return num
}
