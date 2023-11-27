package anagram

import (
	"sort"
	"strings"
)

func mySortString(s string) string {
	sl := strings.Split(s, "")
	sort.Strings(sl)
	return strings.Join(sl, "")
}

func Detect(subject string, candidates []string) []string {
	var sl []string
	subject = strings.ToLower(subject)
	for _, candidate := range candidates {
		candidate_lower := strings.ToLower(candidate)
		if candidate_lower == subject {
			//Do nothing
		} else if mySortString(candidate_lower) == mySortString(subject) {
			sl = append(sl, candidate)
		}
	}
	return sl
}
