package ar

func contains[T comparable](items []T, target T) bool {
	for _, item := range(items) {
		if item == target {
			return true
		}
	}
	return false
}
