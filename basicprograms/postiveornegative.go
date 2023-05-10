package basicprograms

import (
	"errors"
)

func PositiveOrNegative1(num int) (string, error) {
	if num > 0 {
		return "Positve", nil
	} else if num < 0 {
		return "negative", nil
	}
	return "", errors.New("invalid input")
}

func PositiveOrNegative2(num int) (string, error) {
	if num >= 0 {
		if num == 0 {
			return "zero", nil
		} else {
			return "Positve", nil
		}
	} else if num < 0 {
		return "negative", nil
	}
	return "", errors.New("invaild Input")
}
