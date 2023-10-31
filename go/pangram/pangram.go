package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	var alphabet = make(map[string]int)
	for _, s := range strings.ToLower(input) {
		alphabet[string(s)] = 1
	}
	for i := 'a'; i <= 'z'; i++ {
		_, ok := alphabet[string(i)]
		if !ok {
			return false
		}
	}
	return true
}
