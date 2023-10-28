package reverse

func Reverse(input string) string {
	s := make([]rune, len(input))
	//len here is number of bytes, not rune
	if len(input) == 0 {
		return ""
	}

	// need to calculate len in rune here
	n := 0
	for _, v := range input {
		s[n] = v
		//use index by n, not i, because i is where a rune starts. In case of multi bytes, it may not just increase by 1
		//This trick can help us to both initialize the slice and calculate number of rune well
		n++
	}
	s = s[0:n]
	//Need a slice of n runes only

	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
	return string(s)

}
