package basicprograms

func SumOfDigits1(n int) int {
	sum := 0
	for n != 0 {
		rem := n % 10
		sum += rem
		n /= 10
	}
	return sum
}

func SumOfDigits2(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
