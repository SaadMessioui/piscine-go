package piscine

func UltimateDivMod(a *int, b *int) {
	store := *a
	*a = *a / *b
	*b = store % *b
}
