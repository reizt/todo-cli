package utils

func NilSafe[T any](pointer *T) T {
	if pointer == nil {
		return *new(T)
	}
	return *pointer
}
