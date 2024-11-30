package piscine

func UltimateDivMod(a *int, b *int) {
	a1 := *a
	*a = *a / *b
	*b = a1 % *b
}
