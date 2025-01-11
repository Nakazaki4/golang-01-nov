package piscine

func StrRev(s string) string {
	runes := []rune(s)
	p2 := len(runes) - 1

	for i := 0; i < p2; i++ {
		runes[i], runes[p2] = runes[p2], runes[i]
		p2--
	}
	return string(runes)
}
