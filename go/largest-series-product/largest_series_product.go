package lsproduct

import (
	"errors"
	"regexp"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var max = 0
	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	} else if span < 0 {
		return 0, errors.New("span must be positive number")
	}

	r := regexp.MustCompile(`^\d+$`)
	if !r.MatchString(digits) {
		return 0, errors.New("digits input must only contain digits")
	}

	for i := 0; i <= (len(digits) - span); i++ {
		products := 1
		for _, v := range digits[i : i+span] {
			products = products * (int(v) - 48)
		}
		if max <= products {
			max = products
		}
	}
	return int64(max), nil
}
