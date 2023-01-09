package utils

func Includes[T comparable](array []T, value T) int {
	if len(array) == 0 {
		return -1
	}

	for i, v := range array {
		if v == value {
			return i
		}
	}

	return -1
}
