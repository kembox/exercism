package isbn

import (
	"strconv"
	"strings"
)

func isdigit(r rune) bool {
	return (r >= '0' && r <= '9')
}
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	for i, c := range isbn {
		if !isdigit(c) {
		//Or just unicode.IsDidit :|  
		//I've reinvented the wheel

			if i == 9 && c == 'X' {
				continue
			} else {
				return false
			}
		}
	}

	var sum, multiplier int
	for i := 0; i < 10; i++ {
		if i == 9 && string(isbn[i]) == "X" {
			multiplier = 10
		} else {
			multiplier, _ = strconv.Atoi(string(isbn[i]))
		}
		sum += multiplier * (10 - i)
	}
	if sum%11 == 0 {
		return true
	} else {
		return false
	}
}
