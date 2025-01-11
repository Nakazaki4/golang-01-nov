package piscine

func Swap(a *int, b *int) {
	box := *a
	*a = *b
	*b = box
}
