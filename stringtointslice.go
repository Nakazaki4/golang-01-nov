package piscine

func StringToIntSlice(str string) []int {
	slice := []int{}

	if len(str) <= 0 {
		return nil
	}

	for _, v := range str {
		slice = append(slice, int(v))
	}
	return slice
}
