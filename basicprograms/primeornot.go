package basicprograms

import (
	"fmt"
)

func IsPrime1(n int) string {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	if count == 2 {
		return fmt.Sprintf("Prime no : %d", n)
	}
	return fmt.Sprintf("Not Prime : %d", n)
}

func IsPrime2(n int) bool {
	isPrime := true

	if n < 2 {
		return false
	} else {
		for i := 2; i < n; i++ {
			if n%2 == 0 {
				isPrime = false
				break
			}
		}
	}
	return isPrime
}
