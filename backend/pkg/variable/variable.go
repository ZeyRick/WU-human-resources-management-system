package variable

func Create[T any](data T) *T {
	return &data
}

func IntToAlphabet(n int) string {
	return string(n+65)
}
