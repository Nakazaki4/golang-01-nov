package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	words := []string{}
	receipt := make(map[string]int)
	word := ""
	for _, v := range str {
		if v != ' ' {
			word += string(v)
		} else {
			words = append(words, word)
			word = ""
		}
	}
	words = append(words, word)

	for _, item := range words {
		receipt[item]++
	}

	return receipt
}
