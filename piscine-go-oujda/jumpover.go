package piscine

func JumpOver(str string) string {
	runes := []rune(str)
	res := ""

	for i := 2; i < len(runes); i += 3 {
		res += string(runes[i])
	}

	return res + "\n"
}
