package ArrayUtils

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

	return array[from:to]
}
