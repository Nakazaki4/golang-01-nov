package piscine

func Split(s, sep string) []string {
	var words []string
	sepSize := len(sep)
	word := ""

	for index := 0; index < len(s); index++ {
		if s[index] == sep[0] && s[index:index+sepSize] == sep {
			words = append(words, word)
			word = ""
			index = index + sepSize - 1
		} else {
			word = word + string(s[index])
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}
