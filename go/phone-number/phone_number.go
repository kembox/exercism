package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

func Number(phoneNumber string) (string, error) {
	r := regexp.MustCompile(`[^0-9]`)
	var err error
	numbers := r.ReplaceAllString(phoneNumber, "")
	if numbers == "" {
		err = errors.New("invalid input")
	}

	if len(numbers) == 11 {
		if string(numbers[0]) != "1" {
			err = errors.New("invalid input")
		} else {
			numbers = numbers[1:]
		}
	}

	rn := regexp.MustCompile(`^[2-9][0-9]{2}[2-9][0-9]{2}[0-9]{4}$`)
	if !rn.MatchString(numbers) {
		return "", errors.New("invalid input")
	}

	return numbers, err
}

func AreaCode(phoneNumber string) (string, error) {
	numbers, err := Number(phoneNumber)
	area_code := ""
	if err != nil {
		return "", err
	} else {
		area_code = numbers[0:3]
	}
	return area_code, err
}

func Format(phoneNumber string) (string, error) {
	numbers, err := Number(phoneNumber)
	ret := ""
	if err != nil {
		return "", err
	} else {
		ret = fmt.Sprintf("(%s) %s-%s", numbers[:3], numbers[3:6], numbers[6:])
	}
	return ret, nil
}
