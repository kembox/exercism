package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	count := 0
	var err error
	if n == 0 {
		err = errors.New("zero value")
	} else if n < 0 {
		err = errors.New("less than zero value")
	}
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		count += 1
	}
	return count, err
}
