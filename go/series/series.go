package series

func All(n int, s string) []string {
	var sl []string
	for i := 0; i+n <= len(s); i++ {
		sl = append(sl, s[i:i+n])
	}
	return sl
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return s, false
	} else {
		return s[:n], true
	}
}
