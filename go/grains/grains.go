package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Wrong input value")
	} else {
		return uint64(math.Pow(2, float64(number-1))), nil
	}
}

func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		tmp, _ := Square(i)
		sum += tmp
	}
	return sum
}
