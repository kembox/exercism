package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var err error
	if len(a) != len(b) {
		err = errors.New("invalid input. Strings must have the same length")
		return 0, err
	}
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count += 1
		}
	}
	return count, err
}
