package ternary

// Or returns the first non-nil value, or the second value if both are nil.
func Or[T any](a, b T) T {
	if a != nil {
		return a
	}
	return b
}

// And returns the second value if the first value is non-nil, or nil if the first value is nil.
func And[T any](a, b *T) *T {
	if a == nil {
		return nil
	}
	return b
}

// Ternary returns the first value if the statement is true, or the second value if the statement is false.
func Ternary[T any](statement bool, a, b T) T {
	if statement {
		return a
	}
	return b
}
