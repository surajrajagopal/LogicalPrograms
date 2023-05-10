package basicprograms

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

func Factorial1(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial1(n-1)
}
