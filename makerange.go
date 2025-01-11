package piscine

func MakeRange(min, max int) []int {
	if min >= max || (max-min) <= 0 {
		return nil
	}
	size := max - min
	c := make([]int, size)

	for i := 0; i < size; i++ {
		c[i] = i + min
	}
	return c
}
