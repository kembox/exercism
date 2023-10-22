package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function

func DivideFood(fc FodderCalculator, n int) (float64, error) {
	total, err := fc.FodderAmount(n)
	if err == nil {
		factor, ferr := fc.FatteningFactor()
		if ferr == nil {
			return total * factor / float64(n), nil
		} else {
			return 0, ferr
		}
	} else {
		return 0, err
	}
}

// TODO: define the 'ValidateInputAndDivideFood' function

func ValidateInputAndDivideFood(fc FodderCalculator, n int) (float64, error) {
	if n > 0 {
		return DivideFood(fc, n)
	} else {
		return 0, errors.New("invalid number of cows")
	}
}

// TODO: define the 'ValidateNumberOfCows' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

type InvalidCowsError struct {
	NumberOfCows  int
	CustomMessage string
}

func (MyErr *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", MyErr.NumberOfCows, MyErr.CustomMessage)
}

func ValidateNumberOfCows(n int) error {
	if n < 0 {
		return &InvalidCowsError{
			n,
			"there are no negative cows",
		}
	} else if n == 0 {
		return &InvalidCowsError{
			0,
			"no cows don't need food",
		}
	} else {
		return nil
	}
}
