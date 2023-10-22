package scrabble

import (
	"strings"
)

func Score(word string) int {
	points := map[string]int{"AEIOULNRST": 1, "DG": 2, "BCMP": 3, "FHVWY": 4, "K": 5, "JX": 8, "QZ": 10}
	var res int
	for _, s := range word {
		for k, v := range points {
			if strings.Contains(k, strings.ToUpper(string(s))) {
				res += v
			}
		}
	}
	return res
}
