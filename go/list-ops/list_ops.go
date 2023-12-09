package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}
	accu := fn(initial, s[0])
	s[0] = accu
	for i := 0; i < len(s)-1; i += 1 {
		accu = fn(s[i], s[i+1])
		s[i+1] = accu
	}
	return accu
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}
	accu := fn(s[len(s)-1], initial)
	s[len(s)-1] = accu
	for i := len(s) - 1; i > 0; i -= 1 {
		accu = fn(s[i-1], s[i])
		s[i-1] = accu
	}
	return accu

}

func (s IntList) Filter(fn func(int) bool) IntList {
	if len(s) == 0 {
		return s
	}
	var il IntList
	for _, v := range s {
		if fn(v) {
			il = append(il, v)
		}
	}
	return il
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	for i, v := range s {
		s[i] = fn(v)
	}
	return s
}

func (s IntList) Reverse() IntList {
	if len(s) == 0 {
		return s
	}
	l := len(s)
	m := l / 2
	for i := 0; i <= m-1; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, v := range lists {
		s = s.Append(v)
	}
	return s
}
