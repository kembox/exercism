package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	my_map := make(map[string]int, len(word))
	for _, v := range word {
		my_map[strings.ToLower(string(v))] += 1
		if my_map[strings.ToLower(string(v))] > 1 && !strings.Contains(" -", string(v)) {
			return false
		}
	}
	return true
}
