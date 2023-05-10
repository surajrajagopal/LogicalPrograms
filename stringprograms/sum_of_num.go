package stringprograms

import (
	"fmt"
	"strconv"
)

func SumOfNum(str string) (int, error) {
	var sum int = 0
	for _, v := range str {
		if v >= 48 && v <= 57 {
			val, err := strconv.Atoi(string(v))
			if err != nil {
				return 0, fmt.Errorf("%w", err)
			}
			sum = sum + val
		}
	}
	return sum, nil
}
