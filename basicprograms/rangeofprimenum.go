package basicprograms

import "fmt"

func IsPrime(n int) bool {
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
func RangePrime(lower, upper int) {
	for i := lower; i <= upper; i++ {
		if IsPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
}
