package piscine

func Abort(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers[2]
}
