package snacks

func Filter[E any](arr []E, p func(E) bool) []E {
	result := make([]E, 0)

	for _, elt := range arr {
		if p(elt) {
			result = append(result, elt)
		}
	}

	return result
}

func AllOf[E any](arr []E, p func(E) bool) bool {
	for _, elt := range arr {
		if !p(elt) {
			return false
		}
	}

	return true
}

func Count[E any](arr []E, p func(E) bool) int {
	var result int

	for _, elt := range arr {
		if p(elt) {
			result += 1
		}
	}

	return result
}

func Map[E any, O any](arr []E, f func(E) O) []O {
	result := make([]O, len(arr))

	for i, elt := range arr {
		result[i] = f(elt)
	}

	return result
}

func Last[E any](arr []E) E {
	return arr[len(arr)-1]
}
