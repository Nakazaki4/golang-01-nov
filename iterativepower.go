package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	result := nb
	for i := 1; i < power; i++ {
		result = result * nb
	}
	return result
}
