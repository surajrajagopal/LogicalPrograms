package leetcodeprograms

import (
	"strconv"
)

func FizzBuzz(num int) []string {
	var storeFizzBuzz []string
	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			storeFizzBuzz = append(storeFizzBuzz, "FizzBuzz")
		} else if i%3 == 0 {
			storeFizzBuzz = append(storeFizzBuzz, "Fizz")
		} else if i%5 == 0 {
			storeFizzBuzz = append(storeFizzBuzz, "Buzz")
		} else {
			storeFizzBuzz = append(storeFizzBuzz, strconv.Itoa(i))
		}
	}
	return storeFizzBuzz
}
