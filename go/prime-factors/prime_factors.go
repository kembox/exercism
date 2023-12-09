package prime_factors

func Factors(n int64) []int64 {
	if n < 2 {
		return []int64{}
	}
	var ret_sl []int64
	var i int64
	for i = 2; i <= n; i++ {
		for n%i == 0 {
			if n == 1 {
				break
			}
			n = n / i
			ret_sl = append(ret_sl, i)
		}
	}
	return ret_sl
}
