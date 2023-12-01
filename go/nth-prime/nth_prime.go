package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("invalid input")
	}
	cnt := 0
	nth_prime := 0
	for i := 2; cnt < n; i = i + 1 {
		if is_prime(i) {
			cnt += 1
			nth_prime = i
		}
	}
	return nth_prime, nil
}

func is_prime(n int) bool {
	if n == 2 {
		return true
	} else {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
	}
	return true
}
