package piscine

func Join(strs []string, sep string) string {
	c := ""
	for i := 0; i < len(strs); i++ {
		c = c + strs[i]
		if i < len(strs)-1 {
			c += sep
		}
	}
	return c
}
