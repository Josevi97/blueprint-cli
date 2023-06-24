package StringUtils

import (
	"strings"
)

func Join(array []string, join string) string {
	return strings.Join(array, join)
}

func SublistFrom(array []string, from int) []string {
	if array == nil || len(array) < from {
		return []string{}
	}

	return array[from:]
}

func Sublist(array []string, from int, to int) []string {
	if array == nil || len(array) < from {
		return []string{}
	}

	// if from > to {
	// 		// should panic
	// }

	return array[from:to]
}

func GetDataFromMap(data map[string]string) ([]string, []string) {
	keys := make([]string, len(data))
	values := make([]string, len(data))
	i := 0

	for key, value := range data {
		keys[i] = key
		values[i] = value
		i++
	}

	return keys, values
}

func Map(from []string, to string) []string {
	array := make([]string, len(from))
	i := 0

	for i < len(from) {
		array[i] = to
		i++
	}

	return array
}
