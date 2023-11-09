package romannumerals

import (
	"errors"
	"sort"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	var s string
	if input > 3999 || input < 1 {
		return "", errors.New("Invalid input")
	}
	var symbols map[int]string = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	keys := make([]int, 0)
	for k := range symbols {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, v := range keys {
		s += strings.Repeat(symbols[v], input/v)
		input -= (input / v) * v
	}
	return s, nil
}
