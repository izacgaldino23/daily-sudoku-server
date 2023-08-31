package utils

func Map[T comparable, R any](elements []T, callback func(T) R) []R {
	res := make([]R, 0)

	for _, v := range elements {
		res = append(res, callback(v))
	}

	return res
}

func Has[T comparable](elements []T, searched T) bool {
	for i := range elements {
		if elements[i] == searched {
			return true
		}
	}

	return false
}
