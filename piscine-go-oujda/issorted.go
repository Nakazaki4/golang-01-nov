package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	b := true
	c := true
	for i := 0; i < len(a)-1; i++ {
		comp := f(a[i], a[i+1])
		if comp > 0 {
			b = false
		} else {
			if comp < 0 {
				c = false
			}
		}
	}
	return b || c
}
