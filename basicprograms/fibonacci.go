package basicprograms

import "fmt"

func FibonacciSeries(n int) {
	num1, num2 := 0, 1
	for i := 2; i <= n; i++ {
		nextterm := num1 + num2
		num1 = num2
		num2 = nextterm
		fmt.Printf("%d ", num2)
	}
}
