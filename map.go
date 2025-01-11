package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, 0)
	for _, v := range a {
		result = append(result, f(v))
	}
	return result
}
