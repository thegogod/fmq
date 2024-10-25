package slices

func Map[T any, R any](s []T, predicate func(v T) R) []R {
	arr := make([]R, len(s))

	for i, v := range s {
		arr[i] = predicate(v)
	}

	return arr
}
