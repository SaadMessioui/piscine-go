package piscine

func Swap(a *int, b *int) {
	storage := *a
		*a = *b
		*b = storage
}