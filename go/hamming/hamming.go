package hamming

import "errors"

// Distance returns the hamming distance of two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("hamming: strings are of different length")
	}

	distance := 0

	for pos := 0; pos < len(a); pos++ {
		if a[pos] != b[pos] {
			distance++
		}
	}
	return distance, nil
}
