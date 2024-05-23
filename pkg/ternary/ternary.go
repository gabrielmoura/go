package ternary

func Or[T any](a, b T) T {
	if a != nil {
		return a
	}
	return b
}

func And(a, b interface{}) interface{} {
	if a == nil {
		return nil
	}
	return b
}

func Ternary[T any](statement bool, a, b T) T {
	if statement {
		return a
	}
	return b
}
