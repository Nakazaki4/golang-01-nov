package piscine

func ReverseMenuIndex(menu []string) []string {
	res := make([]string, len(menu))
	p1 := 0

	for i := len(menu) - 1; i >= 0; i-- {
		res[p1] = menu[i]
		p1++
	}
	return res
}
