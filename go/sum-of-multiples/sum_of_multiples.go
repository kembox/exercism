package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var m = make(map[int]int)
	for _, v := range divisors {
		if v == 0 {
			break
		}
		for i := v; i < limit; i += v {
			m[i] = 1
		}
	}

	sum := 0
	for k, v := range m {
		if v >= 1 {
			sum += k
		}
	}
	return sum
}
