package piscine

func AppendRange(min, max int) []int {
	var s []int

	for i := min; i < max; i++ {
		s = append(s, i)
	}
	return s
}
