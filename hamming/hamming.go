package hamming

import "errors"

func Distance(a, b string) (int, error) {
	result := 0

	if len(a) != len(b) {
		return result, errors.New("Array lens don't match")
	}

	for i := range a {
		if a[i] != b[i] {
			result++
		}
	}

	return result, nil
}
