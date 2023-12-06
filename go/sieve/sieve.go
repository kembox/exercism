package sieve

func Sieve(limit int) []int {
	composites := make([]bool, limit+1)
	primes := make([]int, limit/2)
	pidx := 0

	for i := 2; i < limit+1; i++ {
		if composites[i] {
			continue
		}

		primes[pidx] = i
		pidx += 1

		for j := 2 * i; j < limit+1; j += i {
			composites[j] = true
		}
	}
	return primes[:pidx]
}
