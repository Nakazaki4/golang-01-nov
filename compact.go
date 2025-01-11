package piscine

func Compact(ptr *[]string) int {
	c := []string{}

	for _, v := range *ptr {
		if v != "" {
			c = append(c, v)
		}
	}

	*ptr = c
	return len(*ptr)
}
