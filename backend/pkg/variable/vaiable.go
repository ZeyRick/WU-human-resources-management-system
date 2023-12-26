package variable

func Create[T any](data T) *T {
	return &data
}