package sieve

func Sieve(limit int) []int {
	var sl []int
	for i := 2; i <= limit; i++ {
		sl = append(sl, i)
	}

	// Avoid risky action that may change origin slice
	// So better create a new slice and append instead of removing item from origin slice
	var composite []int
	for n, v := range sl {
		for _, jv := range sl[n+1:] {
			if jv%v == 0 {
				composite = append(composite, jv)
			}
		}
	}

	var ret []int
	for _, v1 := range sl {
		is_prime := 1
		for _, v2 := range composite {
			if v1 == v2 {
				is_prime = 0
				break
			}
		}
		if is_prime == 1 {
			ret = append(ret, v1)
		}
	}
	return ret
}
