package piscine

func Unmatch(a []int) int {
	count := 0
	for _, v := range a {
		for i := 0; i < len(a); i++ {
			if v == a[i] {
				count++
			}
		}
		if count%2 != 0 {
			return v
		}
	}
	return -1
}
