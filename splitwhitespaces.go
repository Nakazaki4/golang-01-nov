package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	word := ""

	for _, c := range s {
		if c != ' ' && c != '\t' && c != '\n' {
			word = word + string(c)
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}
