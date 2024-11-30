package piscine

func SortIntegerTable(table []int) {
	right := len(table)

	for i := 0; i < right-1; i++ {
		for j := 0; j < right-i-1; j++ {
			if table[j] > table[j+1] {
				temp := table[j]
				table[j] = table[j+1]
				table[j+1] = temp
			}
		}
	}
}
