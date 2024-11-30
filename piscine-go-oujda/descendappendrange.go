package piscine

func DescendAppendRange(max, min int) []int {
	nums := []int{}
	if max <= min {
		return nums
	}
	for i := max; i > min; i-- {
		nums = append(nums, i)
	}
	return nums
}
