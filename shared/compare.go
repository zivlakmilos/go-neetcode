package shared

func CompareArrays[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	m1 := map[T]int{}

	for _, c := range a {
		m1[c]++
	}

	for _, c := range b {
		m1[c]--
	}

	for _, val := range m1 {
		if val != 0 {
			return false
		}
	}

	return true
}

func CompareArrayOfArrays[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		found := false
		for j := range b {
			if CompareArrays(a[i], b[j]) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
