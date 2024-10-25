package slices

func Filter[T any](s []T, predicate func(v T) bool) []T {
	arr := []T{}

	for _, v := range s {
		if predicate(v) {
			arr = append(arr, v)
		}
	}

	return arr
}
