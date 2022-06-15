package commons

func Swap[T Comparable](a, b *T) {
	tmp := *a
	*a = *b
	*b = tmp
}
