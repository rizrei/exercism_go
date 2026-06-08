package hamming

import "errors"

func Distance(a, b string) (int, error) {
	lenA := len(a)
	lenB := len(b)

	if lenA != lenB {
		return 0, errors.New("strands must be of equal length")
	}

	count := 0
	for i := range lenA {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
