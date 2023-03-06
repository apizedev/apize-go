package apize

func Ptr[T any](x T) *T {
	return &x
}

func Int(x int) *int {
	return &x
}

func Bool(x bool) *bool {
	return &x
}
